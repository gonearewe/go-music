package player

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/gonearewe/go-music/library"
)

const (
	RandomMode PlayerMode = iota
	RepeatMode
	SequentialMode
)

const (
	RequestRandomMode Request = iota
	RequestRepeatMode
	RequestSequentialMode
	RequestNextTrack
	RequestPrevTrack
	RequestStop
)

type PlayerMode int

type Request int

type Player struct {
	library     *library.Library
	mode        PlayerMode
	playlist    *playList
	cancel      context.CancelFunc
	handle      *exec.Cmd // handle of the process playing track
	isPlaying   bool
	requestChan chan Request // signal for panel controling
}

func NewPlayer(lib *library.Library) *Player {
	var p Player
	p.library = lib
	p.playlist = newPlayList()
	// by default
	// p.mode= RandomMode
	// p.isPlaying= false
	return &p
}

func (p *Player) Start() chan<- Request {
	var requestChan = make(chan Request, 4)
	// WARNING: miss this statement, you will be blocked forever when filling p.requestChan
	p.requestChan = requestChan
	go func() {
		defer close(requestChan)
		p.workLoop(requestChan)
	}()
	return requestChan
}

func (p *Player) workLoop(requestChan <-chan Request) {
	var request Request
	for {
		select {
		case request = <-requestChan:
		}

		switch request {
		// sets up Player mode of a Player among enumeration.
		case RequestRandomMode:
			p.mode = RandomMode
		case RequestRepeatMode:
			p.mode = RepeatMode
		case RequestSequentialMode:
			p.mode = SequentialMode
		case RequestStop:
			p.stopPlaying()
		case RequestPrevTrack:
			p.stopPlaying()
			p.playlist.pop()
			if addr, ok := p.currentTrackAddr(); !ok {
				continue
			} else {
				p.isPlaying = true // put this outside the function body
				// when the routine exits, isPlaying resets false,
				// which means isPlaying indicates the existence of this routine
				go p.play(addr)
			}
		case RequestNextTrack:
			p.stopPlaying()
			currentTrack, id, _ := p.playlist.peek() // may be nil
			p.playlist.push(p.determineNextTrack(currentTrack, id))
			if addr, ok := p.currentTrackAddr(); !ok {
				continue
			} else {
				go p.play(addr)
			}
		}
	}
}

// determineNextTrack determines next track it will play according to playermode and current track.
func (p *Player) determineNextTrack(cur library.Track, cur_id int) (track library.Track, id int) {
	rand.Seed(time.Now().Unix()) // a seed is required for random

	switch p.mode {
	case RandomMode:
		id = rand.Intn(p.library.NumTracks())
		track = *p.library.GetTrackByID(id)
		// maybe don't want to hear prev song again if possible
		if cur != nil && track == cur && p.library.NumTracks() > 2 {
			return p.determineNextTrack(cur, cur_id)
		}
		return

	case RepeatMode:
		if cur == nil { // when we play it for the first time
			id = rand.Intn(p.library.NumTracks())
			track = *p.library.GetTrackByID(id)
			return
		} else { // replay
			return cur, cur_id
		}

	case SequentialMode:
		if cur == nil { // when we play it for the first time
			id = rand.Intn(p.library.NumTracks())
			track = *p.library.GetTrackByID(id)
			return
		} else {
			id = (cur_id + 1) % p.library.NumTracks() // in case out of boundary
			track = *p.library.GetTrackByID(id)
			return
		}
	}

	panic(fmt.Sprintf("unhandled Player mode: enumberation value %d", p.mode))
}

// currentTrackAddr returns complete address of current track file(not necessarily playing).
func (p *Player) currentTrackAddr() (string, bool) {
	if track, _, err := p.playlist.peek(); err != nil {
		return "", false
	} else {
		return track.FileAddr(), true
	}
}

// // HandleExited tells if the backend process actually playing tracks has exited.
// func (p *Player) HandleExited() bool {
// 	if p.handle == nil || p.handle.ProcessState != nil {
// 		return true
// 	}

// 	return false
// }
