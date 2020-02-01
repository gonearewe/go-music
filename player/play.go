package player

import (
	"context"
	"os"
	"os/exec"

	"github.com/gonearewe/go-music/request"
)

/*
This file contains methods wrapping basic Player methods with panel functions.
*/

// play executes a process blockingly playing tracks and 
// kills the process when receiving anything from done.
func (p *Player) play(done <-chan struct{}, trackAddr <-chan string) {
	var cancel context.CancelFunc
	var state *os.ProcessState

	for {
		select {
		case <-done:
			if cancel!=nil{ // check first to avoid nil-pointer exception
				cancel()
			}
		case addr := <-trackAddr:
			// NOTICE: playing new track requires killing existing track-playing process
			if cancel!=nil{
				cancel()
			}
			var ctx context.Context
			ctx, cancel = context.WithCancel(context.TODO())
			cmd := exec.CommandContext(ctx, "play", addr)
			state = cmd.ProcessState

			go cmd.Run()
		default:
			// a normal exit results in next track
			if state != nil && state.Success() {
				p.requestChan <- request.NewRequestToPlayer(request.RequestNextTrack)
			}
		}
	}
}
