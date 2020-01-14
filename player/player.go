package player

import (
	"math/rand"
	"os/exec"
	"sync"
	"time"

	"github.com/gonearewe/go-music/library"
)

const (
	RandomMode PlayerMode = iota
	RepeatMode
	SequentialMode
)

type PlayerMode int

type Player struct {
	library   *library.Library
	mode      PlayerMode
	status    status
	handle    *exec.Cmd // handle of the process playing track
	locker    *sync.Mutex
	isPlaying bool
	done      chan struct{} // signal for panel controling
}

type status struct {
	prev      *library.Track
	current   *library.Track
	currentID int // only useful in SequentialMode
}

func NewPlayer(lib *library.Library) *Player {
	var p Player
	p.library = lib
	p.locker = new(sync.Mutex)
	// by default
	// p.mode= RandomMode
	// p.isPlaying= false
	return &p
}

func (p *Player) Play() {
	p.checkPreparation()
	p.updateStatus()
	go p.play()
}

// SetMode sets up player mode of a player among enumeration.
func (p *Player) SetMode(mode PlayerMode) {
	p.mode = mode
}

// CurrentTrackAddr returns complete address of current track file(not necessarily playing).
func (p *Player) CurrentTrackAddr() string {
	track := *p.status.current
	return track.FileAddr()
}

// HandleExited tells if the backend process actually playing tracks has exited.
func (p *Player) HandleExited() bool {
	defer p.Unlock()
	p.Lock()

	if p.handle == nil || p.handle.ProcessState != nil {
		return true
	}

	return false
}

func (p *Player) Lock() {
	p.locker.Lock()
}

func (p *Player) Unlock() {
	p.locker.Unlock()
}

// checkPreparation guarantees the player is well initialized.
func (p *Player) checkPreparation() {
	info := "player uninitialized: "

	if p.library == nil {
		panic(info + "empty library")
	}

	if p.isPlaying {
		p.stop()
	}
}

// updateStatus determines current track it will play, and updates prev track for history,
// all according to playermode.
func (p *Player) updateStatus() {
	rand.Seed(time.Now().Unix()) // a seed is required for random

	switch p.mode {
	case RandomMode:
		nextID := rand.Intn(p.library.NumTracks())
		nextTrack := p.library.GetTrackByID(nextID)

		if p.library.NumTracks() > 2 && nextTrack == p.status.current {
			p.updateStatus() // don't want to hear prev song again if possible
			return
		}

		// determine prev track
		// if p.status.prev == nil { // when we play it for the first time
		// 	p.status.prev = p.library.GetTrackByID(rand.Intn(p.library.NumTracks()))
		// }

		// determine current track
		p.status.prev = p.status.current
		p.status.current = nextTrack

	case RepeatMode:
		// determine current track
		if p.status.current == nil { // when we play it for the first time
			p.status.current = p.library.GetTrackByID(rand.Intn(p.library.NumTracks()))
		} else {
			p.status.prev = p.status.current // update history
		}

	case SequentialMode:
		if p.status.current == nil { // when we play it for the first time
			p.status.currentID = 0
			p.status.current = p.library.GetTrackByID(0)
		} else {
			p.status.prev = p.status.current
			if p.status.currentID+1 > p.library.NumTracks()-1 { // reach the last track in the library
				p.status.currentID = 0
				p.status.current = p.library.GetTrackByID(0)
			} else {
				p.status.currentID++
				p.status.current = p.library.GetTrackByID(p.status.currentID)
			}
		}
	}
}

// play executes a process blockingly and records the process in the field 'handle'.
func (p *Player) play() {
	p.Lock()
	cmd := exec.Command("play", p.CurrentTrackAddr())
	p.handle = cmd
	p.isPlaying = true
	p.Unlock()

	cmd.Run()
}

// stop kills the process playing track.
func (p *Player) stop() {
	p.handle.Process.Kill()
	p.handle = nil
	p.isPlaying = false
}
