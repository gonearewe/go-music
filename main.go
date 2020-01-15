package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gonearewe/go-music/config"
	"github.com/gonearewe/go-music/library"
	"github.com/gonearewe/go-music/panel"
	"github.com/gonearewe/go-music/player"
)

var GlobalPlayer = &player.Player{}

func init() {
	// if any errors occur, we won't be able to handle them, it's better to crash.
	defer func() {
		if err := recover(); err != nil {
			TerminateSoftware(err)
		}
	}()

	panel.ShowLOGO()
	// Preparation
	c := &config.LibraryConfiguration{}
	err := config.LoadConfigFromWorkDir(c)
	if err != nil { // no config file found
		GlobalPlayer.SetLibrary()
	} else {
		lib := library.NewLibraryFromConfig(c)
		GlobalPlayer = player.NewPlayer(lib)
	}

	// so that we can see the LOGO
	time.Sleep(2 * time.Second)
}

func main() {
	// TODO: no support for options during to thread safety issues.
	// go GlobalPlayer.ListenForKeyboard()
	for {
		func() {
			// TODO: Oh, lock doesn't solve problems of thread safety, I don't know why.
			defer GlobalPlayer.Unlock()
			GlobalPlayer.Lock()

			if GlobalPlayer.HandleExited() {
				GlobalPlayer.PlayNextTrack()
				for GlobalPlayer.HandleExited() {
				}
			}
		}()
	}
}

// TerminateSoftware terminates the whole program and reports errors.
func TerminateSoftware(info interface{}) {
	panel.EraseScreen()
	fmt.Println("Go Music terminated during to error !!!")
	if err, ok := info.(error); ok {
		fmt.Println(err.Error())
	}
	os.Exit(-1) // TODO: use syscall to kill the whole progress group.
}
