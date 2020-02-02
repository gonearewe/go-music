package player

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gonearewe/go-music/library"
	. "github.com/gonearewe/go-music/request"
)

const (
	RandomMode PlayerMode = iota
	RepeatMode
	SequentialMode
)

type PlayerMode int

type Player struct {
	library  *library.Library
	mode     PlayerMode
	playlist *playList // maintains track history,the one on the top is the one currently playing

	cancel context.CancelFunc // uesd to kill existing track-playing process
	handle chan struct{}      // handle of the process playing track

	// two channels for communication, in and out
	requestChan chan Request   // signal for panel controlling
	outport     chan<- Request // where you send your request to others

	trackCover chan<- string // channel to send cover to panel routine
}

func NewPlayer(lib *library.Library, outport chan<- Request) *Player {
	var p Player
	p.library = lib
	p.playlist = newPlayList()
	p.outport = outport
	p.handle = make(chan struct{}, 4)
	// by default
	// p.mode= RandomMode
	// p.isPlaying= false
	return &p
}

// Start starts a player routine receiving requests and playing tracks accordingly.
//
// text: if player plays next track, it send the track's cover through
// this(to panel routine actually).
//
// it returns a channel through which you may send requests to the player routine.
func (p *Player) Start(text chan<- string) chan<- Request {
	var requestChan = make(chan Request, 4)
	// WARNING: miss this statement, you will be blocked forever when filling p.requestChan
	p.requestChan = requestChan
	p.trackCover = text
	go func() {
		defer close(requestChan)
		p.workLoop(requestChan)
	}()
	return requestChan
}

func (p *Player) workLoop(requestChan <-chan Request) {
	// play is the one actually playing tracks
	var trackAddr = make(chan string, 4)
	go p.play(p.handle, trackAddr)

	var request Request
	for {
		select {
		case request = <-requestChan:
		}

		switch request.Req {
		// sets up Player mode of a Player among enumeration.
		case RequestRandomMode:
			p.mode = RandomMode
		case RequestRepeatMode:
			p.mode = RepeatMode
		case RequestSequentialMode:
			p.mode = SequentialMode

		case RequestStop:
			p.handle <- struct{}{}
			p.outport <- NewRequestToPanel(RequestClearAndStop)

		case RequestPrevTrack:
			p.handle <- struct{}{}
			p.playlist.pop()
			if addr, cover, ok := p.currentTrackInfo(); !ok {
				continue
			} else {
				p.trackCover <- cover // signal panel to change cover
				trackAddr <- addr
			}

		case RequestNextTrack:
			currentTrack, id, _ := p.playlist.peek() // may be nil
			p.playlist.push(p.determineNextTrack(currentTrack, id))
			if addr, cover, ok := p.currentTrackInfo(); !ok {
				continue
			} else {
				p.trackCover <- cover
				trackAddr <- addr
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

// currentTrackInfo returns complete address and cover of current track file(not necessarily playing).
func (p *Player) currentTrackInfo() (addr string, cover string, ok bool) {
	if track, _, err := p.playlist.peek(); err != nil {
		return "", "", false
	} else {
		return track.FileAddr(), track.String(), true
	}
}
