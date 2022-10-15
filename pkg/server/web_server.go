package server

import (
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	"github.com/lambda-honeypot/ccli-tz/pkg/leader"
	"github.com/lambda-honeypot/ccli-tz/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
)

func WebServer(testnetMagic string, cmdRunner leader.CommandRunner, cfg *config.CfgYaml, fileUtils utils.FileUtilsInterface) error {
	http.HandleFunc("/current", func(w http.ResponseWriter, r *http.Request) {
		response := getEpoch(testnetMagic, cmdRunner, 0, cfg, fileUtils)
		_, err := io.WriteString(w, response)
		if err != nil {
			log.Warnf("error in current handler: %v", err)
		}
	})
	http.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		response := getEpoch(testnetMagic, cmdRunner, 1, cfg, fileUtils)
		_, err := io.WriteString(w, response)
		if err != nil {
			log.Warnf("error in next handler: %v", err)
		}
	})
	return http.ListenAndServe(":"+cfg.Port, nil)
}

func readOutputFile(epoch string, utilsInterface utils.FileUtilsInterface) ([]byte, error) {
	dir, err := leader.GetOutputFilePath(utilsInterface)
	if err != nil {
		return []byte{}, err
	}
	path := filepath.Join(dir, epoch+".json")
	return utilsInterface.ReadFile(path)
}

func getEpoch(testnetMagic string, cmdRunner leader.CommandRunner, offset int, cfg *config.CfgYaml, fileUtils utils.FileUtilsInterface) string {
	if !cfg.PersistMode {
		return "persistMode in config is set to false. Please enable for server to function"
	}
	tipData, err := leader.GetTipData(testnetMagic, cmdRunner)
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}
	epoch := strconv.Itoa(tipData.Epoch + offset)
	output, err := readOutputFile(epoch, fileUtils)
	if err != nil {
		return fmt.Sprintf("Failed to read leaderlog file for epoch: %s with below error. Has it been calculated?\n\n%v\n", epoch, err)
	}
	return string(output)
}
