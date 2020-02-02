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
// NOTICE: library.NewLibraryConfiguration() is required during preparation,
// make sure it has already passed the test.
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

	// you should have found the config file "$HOME/.config/go-music/libraries.toml"(as on Linux)
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}

// BenchmarkLoadConfigFromWorkDir includes two benchmarks, one of which also includes the time cost
// of SaveConfigInWorkDir(), comment one and uncomment the other one to choose which one to run benchmark.
// NOTICE: library.NewLibraryConfiguration() and config.SaveConfigInWorkDir() are required during preparation,
// make sure they have already passed the test.
func BenchmarkLoadConfigFromWorkDir(b *testing.B) {
	// Preparation
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = dir + "/../assets" // path to your tracks for testing
	dir = "/media/heathcliff/新加卷/Taylor Swift"

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

	// Benchmark: SaveConfigInWorkDir and LoadConfigFromWorkDir
	b.ResetTimer()
	// I ignore possible returned error here
	for i := 0; i < b.N; i++ {
		config.SaveConfigInWorkDir(libconfig)
		config.LoadConfigFromWorkDir(libconfig)
	}

	// Benchmark: LoadConfigFromWorkDir only
	// config.SaveConfigInWorkDir(libconfig)
	// b.ResetTimer()
	// I ignore possible returned error here
	// for i := 0; i < b.N; i++ {
		// config.LoadConfigFromWorkDir(libconfig)
	// }
}
