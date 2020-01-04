package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// const (
// 	ConfigVersionNum = "0.5.0"
// 	VersionNum       = "0.5.0"
// )

type Configuration interface {
	FileName() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

func GenerateFile(path string, c Configuration) error {
	path = filepath.Clean(path) // config file dir
	if _, err := os.Stat(path); err != nil {
		return errors.New("generate config file: " + err.Error())
	}

	path = filepath.Join(path, c.FileName()) // path including filename

	data, err := c.Marshal()
	if err != nil {
		return errors.New("generate config file: " + err.Error())
	}

	err = ioutil.WriteFile(path, append(fileHeader(), data...), os.ModeSetuid)
	if err != nil {
		return errors.New("generate config file: " + err.Error())
	}

	return nil
}

func fileHeader() []byte {
	return []byte("# This is the config file of go-music\n\n")
}
