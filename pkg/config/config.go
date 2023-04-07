package config

import (
	"github.com/lambda-honeypot/ccli-tz/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const FileName = ".ccli-tz.yaml"

type CfgCreator interface {
	GetConfigFilePath() (string, error)
	ConfigFileExists(filePath string) (bool, error)
	WriteFile(filePath string, yamlData []byte) error
}

type CfgYaml struct {
	VRFSigningKeyFile string `yaml:"VRFSigningKeyFile"`
	StakePoolID       string `yaml:"stakePoolID"`
	GenesisFile       string `yaml:"shelleyGenesisFile"`
	TimeZone          string `yaml:"timeZone"`
	Port              string `yaml:"serverPort"`
	PersistMode       bool   `yaml:"persistMode"`
}

func ReadConfig() *CfgYaml {
	shelleyGenesisFile := utils.NormaliseHomeDir(viper.GetString("shelleyGenesisFile"))
	vrfKeysFile := utils.NormaliseHomeDir(viper.GetString("VRFSigningKeyFile"))
	poolID := viper.GetString("stakePoolID")
	timeZone := viper.GetString("timeZone")
	persistMode := viper.GetBool("persistMode")
	return &CfgYaml{
		VRFSigningKeyFile: vrfKeysFile,
		StakePoolID:       poolID,
		GenesisFile:       shelleyGenesisFile,
		TimeZone:          timeZone,
		PersistMode:       persistMode,
	}
}

func InitialiseConfigFile(creator CfgCreator) {
	filePath, err := creator.GetConfigFilePath()
	if err != nil {
		log.Errorf("failed to get config file path with err: %v", err)
		return
	}
	log.Infof("creating file at: %s", filePath)
	err = writeTemplateConfig(filePath, creator)
	if err != nil {
		log.Errorf("failed to write config file to: %s with err: %v", filePath, err)
	}
}

func writeTemplateConfig(filePath string, creator CfgCreator) error {
	cfgYml := CfgYaml{
		VRFSigningKeyFile: "/path/to/vrf-signing-key-file/vrf.skey",
		StakePoolID:       "<insert pool id>",
		GenesisFile:       "/path/to/genesis-file/shelley-genesis.json",
		TimeZone:          "Europe/London",
		Port:              "9091",
		PersistMode:       true,
	}

	yamlData, err := yaml.Marshal(&cfgYml)

	if err != nil {
		return err
	}
	exists, err := creator.ConfigFileExists(filePath)
	if err != nil {
		return err
	}
	if exists {
		log.Infof("config file exists at: %s, cowardly refusing to init new config", filePath)
		return nil
	}
	err = creator.WriteFile(filePath, yamlData)
	return err
}
