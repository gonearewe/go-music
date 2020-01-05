package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gonearewe/go-music/config"

	"github.com/gonearewe/go-music/library"
)

// Test saving LibraryConfiguration in software's work dir(in user's home dir).
// NOTICE:
func TestSaveConfigInWorkDir(t *testing.T) {
	// Preparation
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = dir + "/../assets" // path to your tracks for testing

	lib, err := library.NewLibrary("MyLibrary", dir)
	err = lib.Scan()
	fmt.Println(lib.NumTracks())
	if err != nil {
		panic(err)
	}

	libconfig := library.NewLibraryConfiguration(lib)

	// Generate filepath for validating.
	path, err := os.UserConfigDir()
	if err != nil {
		panic(err) // panic here since it's not tested function's fault
	}
	path = filepath.Join(path, config.WorkDirName, libconfig.FileName())

	// Test
	config.SaveConfigInWorkDir(libconfig)

	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}
