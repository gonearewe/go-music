package library_test

import (
	"os"
	"testing"

	"github.com/gonearewe/go-music/library"
)

var names = []string{"01", "abc", "_34h", `//we`, `\uvw`, ` ab`, `A B`, `我的库`, `ss&%!^*ss`}
var invalidPaths = []string{"", "/usr/bin/", "~/", "noway", "/tmp"}

func TestScan(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println("current dir: %s", dir)
	aValidPath := dir + "/../assets"
	validPaths := []string{aValidPath, aValidPath + "/./"}

	// initialize library with kinds of name
	for _, name := range names {
		_, err := library.NewLibrary(name, aValidPath)
		if err != nil {
			t.Errorf("initialize library with kinds of name and valid path: %s", err.Error())
		}
	}

	// initialize library with invalid path
	for _, invalidPath := range invalidPaths {
		lib, err := library.NewLibrary("MyLibrary", invalidPath)
		if err == nil { // errors reports expected
			err := lib.Scan()
			if err == nil {
				t.Errorf("initialize library with invalid path: expected error not found: %s", invalidPath)
			}
		}
	}

	// scan library, pass expected
	for _, validPath := range validPaths {
		lib, err := library.NewLibrary("MyLibrary", validPath)
		if err != nil {
			t.Errorf("initialize library with valid params: %s", err.Error())
		}

		err = lib.Scan()
		if err != nil {
			t.Errorf("scan library: %s", err.Error())
		}
	}
}

func TestScanWithRoutines(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println("current dir: %s", dir)
	aValidPath := dir + "/../assets"
	validPaths := []string{aValidPath, aValidPath + "/./"}

	// initialize library with kinds of name
	for _, name := range names {
		_, err := library.NewLibrary(name, aValidPath)
		if err != nil {
			t.Errorf("initialize library with kinds of name and valid path: %s", err.Error())
		}
	}

	// initialize library with invalid path
	for _, invalidPath := range invalidPaths {
		lib, err := library.NewLibrary("MyLibrary", invalidPath)
		if err == nil { // errors reports expected
			err := lib.ScanWithRoutines()
			if err == nil {
				t.Errorf("initialize library with invalid path: expected error not found: %s", invalidPath)
			}
		}
	}

	// scan library, pass expected
	for _, validPath := range validPaths {
		lib, err := library.NewLibrary("MyLibrary", validPath)
		if err != nil {
			t.Errorf("initialize library with valid params: %s", err.Error())
		}

		err = lib.ScanWithRoutines()
		if err != nil {
			t.Errorf("scan library: %s", err.Error())
		}
	}
}

// BenchmarkScan requires a folder containing large numbers of tracks,
// I use a folder on my computer and you may select one on yours.
func BenchmarkScan(b *testing.B) {
	// Preparation
	const dir = "/media/heathcliff/新加卷/LOSSLESS MUSIC" // path to your tracks for testing

	lib, err := library.NewLibrary("MyLibrary", dir)
	if err != nil {
		b.Errorf("initialize library with valid params: %s", err.Error())
	}

	// Benckmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lib.Scan()
	}
}

// BenchmarkScanWithRoutines is about 50% faster than BenchmarkScan.
//
// Also, it requires a folder containing large numbers of tracks,
// I use a folder on my computer and you may select one on yours.
func BenchmarkScanWithRoutines(b *testing.B) {
	// Preparation
	const dir = "/media/heathcliff/新加卷/LOSSLESS MUSIC" // path to your tracks for testing

	lib, err := library.NewLibrary("MyLibrary", dir)
	if err != nil {
		b.Errorf("initialize library with valid params: %s", err.Error())
	}

	// Benckmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lib.ScanWithRoutines()
	}
}
