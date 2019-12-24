package player

import (
	"math/rand"
	"os/exec"

	"github.com/gonearewe/go-music/library"
)

const (
	randomMode playerMode = iota
	repeatMode
	sequentialMode
)

type playerMode = int

type Player struct {
	library   *library.Library
	mode      playerMode
	status    status
	handle    *exec.Cmd // handle of the process playing track
	isPlaying bool
}

type status struct {
	prev      *library.Track
	current   *library.Track
	currentID int // only useful in sequentialMode
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

// updateStatus determines current track it will play, and updates prev track for history,
// all according to playermode.
func (p *Player) updateStatus() {
	switch p.mode {
	case randomMode:
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

	case repeatMode:
		// determine current track
		if p.status.current == nil { // when we play it for the first time
			p.status.current = p.library.GetTrackByID(rand.Intn(p.library.NumTracks()))
		} else {
			p.status.prev = p.status.current // update history
		}

	case sequentialMode:
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

// play executes a process non-blockingly and records the process in the field 'handle'.
func (p *Player) play() {
	cmd := exec.Command("play", p.CurrentTrackAddr())
	cmd.Start()
	p.handle = cmd
	// cmd.Process.Kill()
}

// stop kills the process playing track.
func (p *Player) stop() {
	p.handle.Process.Kill()
	p.handle = nil
}
