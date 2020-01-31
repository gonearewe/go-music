package player

import (
	"context"
	"os/exec"
)

/*
This file contains methods wrapping basic Player methods with panel functions.
*/

// stopPlaying kills the process playing track safely.
func (p *Player) stopPlaying() {
	defer func() { p.isPlaying = false }()
	defer recover()  // in case
	if p.isPlaying { // play in progress
		for p.cancel == nil {
		} // program has been initialized but hasn't run

		// p.cancel() will call p.handle.Process.Kill() if possible
		// but cause no nil-pointer exception if p.handle.Process is nil
		p.cancel()
		p.handle = nil
		p.cancel = nil
	}
}

// play executes a process blockingly and records the process in the field 'handle'.
func (p *Player) play(trackAddr string) {
	// when the routine exits, isPlaying resets false,
	// which means isPlaying indicates the existence of this routine
	p.isPlaying = true
	// WARNING: resets isPlaying before signal next track
	defer func() {
		if p.handle.ProcessState == nil {
			p.requestChan <- RequestNextTrack
		}
	}()
	defer func() { p.isPlaying = false }()
	var ctx context.Context
	ctx, p.cancel = context.WithCancel(context.TODO())
	cmd := exec.CommandContext(ctx, "play", trackAddr)
	p.handle = cmd

	cmd.Run()
}

// func (p *Player) PlayNextTrack() {
// 	if p.done != nil {
// 		// signaling goroutine randering cover to terminate.
// 		close(p.done)
// 	} else {
// 		panel.EraseScreen()
// 	}

// 	p.Play()

// 	p.done = make(chan struct{}, 1)
// 	go p.ShowCover(p.done)
// }

// func (p *Player) PlayPrevTrack() {
// 	defer p.Unlock()
// 	p.Lock()

// 	if p.done != nil {
// 		// signaling goroutine randering cover to terminate.
// 		close(p.done)
// 	} else {
// 		panel.EraseScreen()
// 	}

// 	p.checkPreparation()
// 	p.playPrevTrack()

// 	p.done = make(chan struct{}, 1)
// 	go p.ShowCover(p.done)
// }

// func (p *Player) playPrevTrack() {
// 	if p.status.prev == nil {
// 		return
// 	}

// 	if p.mode == SequentialMode {
// 		if p.status.currentID != 0 {
// 			p.status.currentID--
// 		}

// 		if p.status.currentID == 0 {
// 			p.status.prev = nil
// 		} else {
// 			p.status.prev = p.library.GetTrackByID(p.status.currentID - 1)
// 		}

// 		p.status.current = p.library.GetTrackByID(p.status.currentID)
// 	} else {
// 		// swap prev track with current one
// 		p.status.prev, p.status.current = p.status.current, p.status.prev
// 	}

// 	go p.play()
// }
