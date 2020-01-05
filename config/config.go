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

const WorkDirName = "go-music"

type Configuration interface {
	FileName() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

// LoadConfigFromWorkDir reads config file in software's work dir(in user's home dir)
// and fills given Configuration with unmarshal config.
func LoadConfigFromWorkDir(c Configuration) error {
	path, err := os.UserConfigDir()
	if err != nil {
		return errors.New("load config file from work dir: " + err.Error())
	}
	path = filepath.Join(path, WorkDirName, c.FileName())

	// check path validity
	if _, err = os.Stat(path); err != nil {
		return errors.New("load config file from work dir: " + err.Error())
	}

	// load config
	if data, err := ioutil.ReadFile(path); err != nil {
		return errors.New("load config file from work dir: " + err.Error())
	} else {
		return c.Unmarshal(data)
	}
}

func SaveConfigInWorkDir(c Configuration) error {
	path, err := os.UserConfigDir() // "$HOME/.config" on Linux
	if err != nil {
		return errors.New("save config file in work dir: " + err.Error())
	}
	path = filepath.Join(path, WorkDirName)
	os.Mkdir(path, os.ModePerm) // TODO: The os.ModePerm may give unnessary x permission.
	path = filepath.Join(path, c.FileName())

	if data, err := c.Marshal(); err != nil {
		return errors.New("save config file in work dir: " + err.Error())
	} else {
		// TODO: The os.ModePerm may give unnessary x permission.
		return ioutil.WriteFile(path, append(fileHeader(), data...), os.ModePerm)
	}
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
