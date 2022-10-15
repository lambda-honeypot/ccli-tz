package leader

import (
	"encoding/json"
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	"github.com/lambda-honeypot/ccli-tz/pkg/utils"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type CommandRunner interface {
	RunCardanoCmd(trimmedArgs []string) (string, error)
}

type OutputJSON struct {
	Epoch    int           `json:"epoch"`
	PoolID   string        `json:"poolID"`
	Schedule []ScheduleRow `json:"schedule"`
}
type ScheduleRow struct {
	BlockCount    int
	SlotNumber    int
	ScheduledTime string
}

type ConfigGetter interface {
	GetConfig() config.CfgYaml
}

func CreateAndRun(periodShort string, testnetMagic string, cmdRunner CommandRunner, cfg *config.CfgYaml, utilsInterface utils.FileUtilsInterface) error {
	period := "--" + periodShort
	shelleyGenesisFile := cfg.GenesisFile
	vrfKeysFile := cfg.VRFSigningKeyFile
	poolID := cfg.StakePoolID
	timeZone := cfg.TimeZone
	fmt.Println(fmt.Sprintf("Calculating for pool: %s", poolID))

	tipData, err := GetTipData(testnetMagic, cmdRunner)
	if err != nil {
		return err
	}
	if tipData.SyncProgress != "100.00" {
		return fmt.Errorf("tip not sync'd - please wait until 100.00. Current %s", tipData.SyncProgress)
	}
	epoch := tipData.Epoch
	if periodShort == "next" {
		epoch++
	}
	fmt.Println(fmt.Sprintf("Calculating for epoch: %d", epoch))
	leaderArgs := CalculateLeaderArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
	schedule, err := CalcTZSchedule(timeZone, leaderArgs, cmdRunner)
	outputJSON := OutputJSON{
		Schedule: schedule,
		Epoch:    epoch,
		PoolID:   poolID,
	}
	if err != nil {
		return err
	}
	output := GenerateScheduleOutput(outputJSON, periodShort)
	fmt.Println(string(output))
	if cfg.PersistMode {
		return writeOutputToFile(epoch, output, utilsInterface)
	}
	return nil
}

func LogOutParams(periodShort string, testnetMagic string, cfg *config.CfgYaml) {
	period := "--" + periodShort
	shelleyGenesisFile := cfg.GenesisFile
	vrfKeysFile := cfg.VRFSigningKeyFile
	poolID := cfg.StakePoolID
	trimmedArgs := CalculateLeaderArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
	log.Infof("dry-run, would have executed:\n\ncardano-cli %v", trimmedArgs)
}

func GenerateScheduleOutput(outputJSON OutputJSON, period string) []byte {
	if len(outputJSON.Schedule) == 0 {
		return []byte(fmt.Sprintf("No schedule blocks for %s epoch", period))
	}
	b, err := json.MarshalIndent(outputJSON, "", "  ")
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return b
}

func CalcTZSchedule(timeZone string, trimmedArgs []string, runner CommandRunner) ([]ScheduleRow, error) {
	var rows []ScheduleRow
	rawSchedule, err := runner.RunCardanoCmd(trimmedArgs)
	if err != nil {
		log.Errorf("failed to run leader log with: %s", rawSchedule)
		return rows, err
	}
	lines := strings.Split(rawSchedule, "\n")
	for i, line := range lines[2:] {
		spaceSplit := splitLine(line)
		if len(spaceSplit) != 2 {
			continue
		}
		row := createRow(timeZone, spaceSplit, i)
		rows = append(rows, row)
	}
	return rows, nil
}

func GetOutputFilePath(utilsInterface utils.FileUtilsInterface) (string, error) {
	userPath, err := utilsInterface.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(userPath, ".ccli-files")
	return filePath, nil
}

func writeOutputToFile(epoch int, output []byte, utilsInterface utils.FileUtilsInterface) error {
	dir, err := GetOutputFilePath(utilsInterface)
	if err != nil {
		return err
	}
	err = utilsInterface.MkDir(dir)
	if err != nil {
		return err
	}
	epochStr := strconv.Itoa(epoch)
	path := filepath.Join(dir, epochStr+".json")
	return utilsInterface.WriteFile(path, output)
}

func splitLine(line string) []string {
	rawSpaceSplit := strings.Split(strings.TrimSpace(line), "  ")
	var spaceSplit []string
	for _, elem := range rawSpaceSplit {
		if strings.TrimSpace(elem) != "" {
			spaceSplit = append(spaceSplit, elem)
		}
	}
	return spaceSplit
}

func createRow(timeZone string, spaceSplit []string, i int) ScheduleRow {
	rawTS := strings.TrimSpace(spaceSplit[len(spaceSplit)-1])
	convertedTime, err := convertTime(rawTS, timeZone)
	if err != nil {
		log.Errorf("failed to convert time with: %v", err)
	}
	slot, err := strconv.Atoi(strings.TrimSpace(spaceSplit[0]))
	if err != nil {
		log.Errorf("failed to convert slot num with: %v", err)
	}
	row := ScheduleRow{BlockCount: i + 1, SlotNumber: slot, ScheduledTime: convertedTime}
	return row
}

func convertTime(timeStamp, timeZone string) (string, error) {
	layout := "2006-01-02 15:04:05 MST"
	t, err := time.Parse(layout, timeStamp)
	if err != nil {
		return timeStamp, err
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return timeStamp, err
	}
	return t.In(loc).Format(layout), nil
}
