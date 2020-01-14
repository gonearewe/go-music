package player

import (
	"github.com/gonearewe/go-music/panel"
)

/*
This file contains methods wrapping basic player methods with panel functions.
*/

func (p *Player) PlayNextTrack() {
	if p.done != nil {
		// signaling goroutine randering cover to terminate.
		close(p.done)
	} else {
		panel.EraseScreen()
	}

	p.Play()

	p.done = make(chan struct{}, 1)
	go p.ShowCover(p.done)
}

func (p *Player) PlayPrevTrack() {
	defer p.Unlock()
	p.Lock()

	if p.done != nil {
		// signaling goroutine randering cover to terminate.
		close(p.done)
	} else {
		panel.EraseScreen()
	}

	p.checkPreparation()
	p.playPrevTrack()

	p.done = make(chan struct{}, 1)
	go p.ShowCover(p.done)
}

func (p *Player) playPrevTrack() {
	if p.status.prev == nil {
		return
	}

	if p.mode == SequentialMode {
		if p.status.currentID != 0 {
			p.status.currentID--
		}

		if p.status.currentID == 0 {
			p.status.prev = nil
		} else {
			p.status.prev = p.library.GetTrackByID(p.status.currentID - 1)
		}

		p.status.current = p.library.GetTrackByID(p.status.currentID)
	} else {
		// swap prev track with current one
		p.status.prev, p.status.current = p.status.current, p.status.prev
	}

	go p.play()
}
