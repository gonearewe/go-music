package player_test

import (
	"github.com/gonearewe/go-music/request"
	"time"
	"github.com/gonearewe/go-music/player"
	"os"
	"testing"

	"github.com/gonearewe/go-music/library"
)

// Unlike ordinary tests where you validate output, you will listen to tracks by yourself now.
func TestStart(t *testing.T) {
	// Preparation
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = dir + "/../assets" // path to your tracks for testing

	lib, err := library.NewLibrary("MyLibrary", dir)
	if err != nil {
		t.Errorf("initialize library with valid params: %s", err.Error())
	}

	err = lib.ScanWithRoutines()
	if err != nil {
		t.Errorf("scan library: %s", err.Error())
	}

	// Test
	p:=player.NewPlayer(lib,make(chan request.Request, 1000))

	var requests=[]request.RequestType{
		request.RequestNextTrack,
		request.RequestNextTrack,
		request.RequestNextTrack,
		request.RequestPrevTrack,
		request.RequestPrevTrack,
		request.RequestRepeatMode,
	}
	ch:=p.Start(make(chan string, 1000))
	for _,req:=range requests{
		ch<-request.NewRequestToPlayer(req)
		time.Sleep(3*time.Second)
	}
	time.Sleep(10*time.Second)
	// Another process is started on the backend, this process can exit naturally.
}
