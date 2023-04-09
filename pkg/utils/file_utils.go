package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type FileUtilsInterface interface {
	WriteFile(path string, contents []byte) error
	MkDir(path string) error
	ReadFile(path string) ([]byte, error)
	UserHomeDir() (string, error)
}

type FileUtils struct{}

func (FileUtils) WriteFile(path string, contents []byte) error {
	return os.WriteFile(path, contents, 0666)
}

func (FileUtils) MkDir(path string) error {
	return os.MkdirAll(path, 0777)
}

func (FileUtils) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (FileUtils) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func NormaliseHomeDir(file string) string {
	if strings.HasPrefix(file, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Errorf("failed to get userhome dir with: %v", err)
			return file
		}
		return strings.Replace(file, "~", home, 1)
	}
	return file
}
