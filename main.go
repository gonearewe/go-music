package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gonearewe/go-music/config"
	"github.com/gonearewe/go-music/library"
	"github.com/gonearewe/go-music/panel"
	"github.com/gonearewe/go-music/player"
	. "github.com/gonearewe/go-music/request"
)

var GlobalPlayer = &player.Player{}
var port =make(chan Request, 20)

func init() {
	// if any errors occur, we won't be able to handle them, it's better to crash.
	defer func() {
		if err := recover(); err != nil {
			TerminateSoftware(err)
		}
	}()

	// Preparation
	c := &config.LibraryConfiguration{}
	err := config.LoadConfigFromWorkDir(c)
	if err != nil { // no config file found
		// GlobalPlayer.SetLibrary()
	} else {
		lib := library.NewLibraryFromConfig(c)
		GlobalPlayer = player.NewPlayer(lib,port)
	}

	// so that we can see the LOGO
	time.Sleep(2 * time.Second)
}

func main() {

	// var requests=[]player.Request{R}
	panelRequestChan,coverChan:=panel.Start(make(chan struct{}), port)
	panelRequestChan<-NewRequestToPanel(RequestShowLOGO)
	// time.Sleep(1*time.Second)
	playerRequestChan := GlobalPlayer.Start(coverChan)
	playerRequestChan<-NewRequestToPlayer(RequestNextTrack)
	for {
		select{
		case mail:=<-port:
			switch mail.Destination{
			case PLAYER:
				playerRequestChan<-mail
			case PANEL:
				panelRequestChan<-mail
			}
		}
	}
}

// func main(){
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}
// 	dir = dir + "/./assets" // path to your tracks for testing

// 	lib, err := library.NewLibrary("MyLibrary", dir)
// 	if err != nil {
		
// 	}

// 	err = lib.ScanWithRoutines()
// 	if err != nil {
		
// 	}

// 	// Test
// 	p:=player.NewPlayer(lib,make(chan Request, 1000))

// 	var requests=[]RequestType{
// 		RequestNextTrack,
// 		RequestNextTrack,
// 		RequestNextTrack,
// 		RequestPrevTrack,
// 		RequestPrevTrack,
// 		RequestRepeatMode,
// 	}
// 	ch:=p.Start(make(chan string, 1000))
// 	cnt:=0
// 	for _,req:=range requests{
// 		ch<-NewRequestToPlayer(req)
// 		time.Sleep(3*time.Second)
// 		cnt++
// 		println(cnt)
// 	}
// 	time.Sleep(10*time.Second)
// }

// TerminateSoftware terminates the whole program and reports errors.
func TerminateSoftware(info interface{}) {
	panel.EraseScreen()
	fmt.Println("Go Music terminated during to error !!!")
	if err, ok := info.(error); ok {
		fmt.Println(err.Error())
	}
	os.Exit(-1) // TODO: use syscall to kill the whole progress group.
}
