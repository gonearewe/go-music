package main

import (
	"fmt"
	"os"

	"github.com/gonearewe/go-music/config"
	"github.com/gonearewe/go-music/library"
	"github.com/gonearewe/go-music/panel"
	"github.com/gonearewe/go-music/player"
	. "github.com/gonearewe/go-music/request"
)

var GlobalPlayer = &player.Player{}
var port = make(chan Request, 20)

func main() {
	// if any errors occur, we won't be able to handle them, it's better to crash.
	defer func() {
		if err := recover(); err != nil {
			TerminateSoftware(err)
		}
	}()

	// Init Panel
	panelRequestChan, coverChan := panel.Start(make(chan struct{}), port)
	panelRequestChan <- NewRequestToPanel(RequestShowLOGO)

	// Prepare Library and Player
	c := &config.LibraryConfiguration{}
	err := config.LoadConfigFromWorkDir(c)
	if err != nil { // no config file found
		libName, libPath := panel.RequestNewLibrary()
		if err = library.ScanLibraryAndSaveConfig(libName, libPath); err != nil {
			panic(err)
		}

		os.Exit(0) // reboot to load library
	} else {
		lib := library.NewLibraryFromConfig(c)
		GlobalPlayer = player.NewPlayer(lib, port)
	}

	// Start Player
	playerRequestChan := GlobalPlayer.Start(coverChan)
	playerRequestChan <- NewRequestToPlayer(RequestNextTrack) // fuse to boot
	for {
		select {
		case mail := <-port:
			switch mail.Destination {
			case PLAYER:
				playerRequestChan <- mail
			case PANEL:
				panelRequestChan <- mail
			case MAIN:
				// NOTE: of course, here comes unsafe code, it may cause
				// exceptions, but what we can do apart from ignoring them.
				infos, _ := mail.Attachments.([]string)
				library.ScanLibraryAndSaveConfig(infos[0], infos[1])
			}
		}
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
