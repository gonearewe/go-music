package player

import (
	"math/rand"
	"os/exec"

	"github.com/gonearewe/go-music/library"
)

const (
	randomMode playerMode = iota
	//sequentialMode playerMode
)

type playerMode = int

type Player struct {
	library   *library.Library
	mode      playerMode
	status    status
	isPlaying bool
}

type status struct {
	current *library.Track
	next    *library.Track
}

func (p *Player) Play() {
	p.updateStatus()
	p.play()
}

// CurrentTrackAddr returns complete address of current track file(not necessarily playing).
func (p *Player) CurrentTrackAddr() string {
	track := *p.status.current
	return track.FileAddr()
}

// updateStatus determines next track it will play, and determines current track if not determined,
// all according to playermode.
func (p *Player) updateStatus() {
	if p.mode == randomMode {
		nextID := rand.Intn(p.library.NumTracks())
		nextTrack := p.library.GetTrackByID(nextID)
		if p.library.NumTracks() > 2 && nextTrack == p.status.current {
			p.updateStatus() // don't want to hear current song again if possible
			return
		}

		// determine current track
		if p.status.current == nil { // when we play it for the first time
			p.status.current = p.library.GetTrackByID(rand.Intn(p.library.NumTracks()))
		}

		// determine next track
		p.status.next = nextTrack
	}
}

func (p *Player) play() {
	cmd := exec.Command("play", p.CurrentTrackAddr())
	cmd.Start()
	// cmd.Process.Kill()
}
