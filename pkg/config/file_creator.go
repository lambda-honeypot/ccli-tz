package config

import (
	"os"
	"path/filepath"
)

type FileConfigCreator struct{}

func (FileConfigCreator) GetConfigFilePath() (string, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(userPath, FileName)
	return filePath, nil
}

func (FileConfigCreator) ConfigFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	switch err.(type) {
	case *os.PathError:
		return false, nil
	default:
		return false, err
	}
}

func (FileConfigCreator) WriteFile(filePath string, yamlData []byte) error {
	err := os.WriteFile(filePath, yamlData, 0666)
	return err
}
