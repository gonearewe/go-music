package library_test

import (
	"os"
	"testing"

	"github.com/gonearewe/go-music/library"
)

var names = []string{"", "01", "abc", "_34h", `//we`, `\uvw`, ` ab`, `A B`, `我的库`, `ss&%!^*ss`}
var invalidPaths = []string{"", "/usr/bin/", "~/", "noway", "/tmp"}

func TestScan(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println("current dir: %s", dir)
	aValidPath := dir + "/../assets"
	validPaths:=[]string{aValidPath,aValidPath+"/./"}

	for _, name := range names {
		_, err := library.NewLibrary(name, aValidPath)
		// if lib.name!=name{
		// 	t.Fatalf("initialize library with name: want %s, get: %s", name,lib.name)
		// }
		if err != nil {
			t.Errorf("initialize library with kinds of name and valid path: %s", err.Error())
		}
	}

	for _, invalidPath := range invalidPaths {
		lib, err := library.NewLibrary("MyLibrary", invalidPath)
		if err != nil {
			println("initialize library with invalid path: successfully find error: %s: %s", err.Error(), invalidPath)
		} else {
			err := lib.Scan()
			if err == nil {
				t.Errorf("initialize library with invalid path: expected error not found: %s", invalidPath)
			} else {
				println("initialize library with invalid path: successfully find error: %s: %s", err.Error(), invalidPath)
			}

		}
	}

	for _,validPath:=range validPaths{
		lib, err := library.NewLibrary("MyLibrary", validPath)
		if err != nil {
			t.Errorf("initialize library with valid params: %s", err.Error())
		}

		err = lib.Scan()
		if err != nil {
			t.Errorf("scan library: %s", err.Error())
		}

		println(lib.String())
	}
}
