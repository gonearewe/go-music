package player

import (
	"math/rand"
	"github.com/gonearewe/go-music/library"
)

const (
	randomMode playerMode =iota
	//sequentialMode playerMode
)

type playerMode =int

type Player struct{
	library *library.Library
	mode playerMode
	status  status
}

type status struct{
	current *library.Track
	next *library.Track
}

func (p *Player)Play(){
	p.updateStatus()
}

func (p *Player)updateStatus(){
	if p.mode==randomMode{
		nextID:=rand.Intn(p.library.NumTracks())
	}
}