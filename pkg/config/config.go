package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const FileName = ".ccli-tz"

type CfgCreator interface {
	GetConfigFilePath() (string, error)
	ConfigFileExists(filePath string) (bool, error)
	WriteFile(filePath string, yamlData []byte) error
}

type CfgYaml struct {
	VRFSigningKeyFile string
	StakePoolID       string
	GenesisFile       string
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
