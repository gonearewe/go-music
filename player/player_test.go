package player_test

import (
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

	err = lib.Scan()
	if err != nil {
		t.Errorf("scan library: %s", err.Error())
	}

	// Test
	p:=player.NewPlayer(lib)

	// var requests=[]player.Request{R}
	ch:=p.Start()
	ch<-player.RequestNextTrack
	time.Sleep(10*time.Second)
	// Another process is started on the backend, this process can exit naturally.
}
