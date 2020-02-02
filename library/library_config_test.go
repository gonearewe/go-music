package library_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gonearewe/go-music/config"
	"github.com/gonearewe/go-music/library"
	"github.com/pelletier/go-toml"
)

// TestNewLibraryConfiguration tests if the function can produce correct struct and pass go-toml's transformation.
func TestNewLibraryConfiguration(t *testing.T) {
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

	// Test
	libconfig := library.NewLibraryConfiguration(lib)

	m, _ := toml.Marshal(libconfig) // marshal

	// NOTE: following statement initializes newlibconfig as a nil pointer
	// rather than a pointer to a empty struct.
	// var newlibconfig *config.LibraryConfiguration
	newlibconfig := &config.LibraryConfiguration{} // the correct statement
	toml.Unmarshal(m, newlibconfig)                // unmarshal back

	fmt.Println(*newlibconfig)
}
