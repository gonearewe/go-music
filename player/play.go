package player

import (
	"context"
	"os/exec"

	"github.com/gonearewe/go-music/request"
)

// play executes a process blockingly playing tracks and
// kills the process when receiving anything from done.
func (p *Player) play(done <-chan struct{}, trackAddr <-chan string) {
	var cancel context.CancelFunc
	// var successfullyexited = make(chan struct{}, 1)

	for {
		select {
		case <-done:
			if cancel != nil { // check first to avoid nil-pointer exception
				cancel()
			}
		case addr := <-trackAddr:
			// NOTICE: playing new track requires killing existing track-playing process
			if cancel != nil {
				cancel()
			}
			var ctx context.Context
			ctx, cancel = context.WithCancel(context.TODO())
			cmd := exec.CommandContext(ctx, "play", addr)

			go func() {
				cmd.Run()
				if cmd.ProcessState.Success() {
					p.requestChan <- request.NewRequestToPlayer(request.RequestNextTrack)
				}
			}()

			// a normal exit results in next track
			// case <-successfullyexited:
			// 	p.requestChan <- request.NewRequestToPlayer(request.RequestNextTrack)
		}
	}
}
