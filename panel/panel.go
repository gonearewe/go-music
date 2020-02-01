package panel

import (
	"fmt"
	"os"
	"time"

	. "github.com/gonearewe/go-music/request"
)

const (
	LOGO = `
  ____         __  __           _      
 / ___| ___   |  \/  |_   _ ___(_) ___ 
| |  _ / _ \  | |\/| | | | / __| |/ __|
| |_| | (_) | | |  | | |_| \__ \ | (__ 
 \____|\___/  |_|  |_|\__,_|___/_|\___|
                                       
`
)

func showLOGO() {
	EraseScreen()
	fmt.Println(LOGO)
}

func Start(done <-chan struct{}, outport chan<- Request) (requests chan<- Request,
	trackInfos chan<- string) {

	var req = make(chan Request, 4)
	var info = make(chan string, 4)
	go func() {
		defer close(req)
		defer close(info)
		panelLoop(done, req, info, outport)
	}()
	return req, info
}

func panelLoop(done <-chan struct{}, requests <-chan Request,
	trackInfos <-chan string, outport chan<- Request) {

	var cover = ""
	var theme ColorTheme

	var buf = make([]byte, 1)
	go func(){os.Stdin.Read(buf)}()

	for {
		select {
		case <-done:
			EraseScreen()
			return

		case request := <-requests:
			if request.Req == RequestShowLOGO {
				showLOGO()
				time.Sleep(2 * time.Second) // so that we can see the LOGO
			} else if request.Req == RequestStop {
				EraseScreen()
				cover = "" // no more refresh
			}

		case cover = <-trackInfos:
			// TODO: we can't choose color theme by ourself.
			theme = RandomColorTheme()
			cover = RenderText(cover, theme)

		default:
		}

		if buf[0]!=0 {
			var input = string(buf)
			buf[0]=0
			listenForKeyboard(input, outport)

			go func(){os.Stdin.Read(buf)}()
		}

		if cover != "" {
			showCover(cover, theme)
		}
	}
}

// showCover should be executed periodically to refresh terminal image.
func showCover(text string, theme ColorTheme) {
	const ProgressBarRefreshIntervalMs = 10 // how much time(ms) progressbar is refreshed
	EraseScreen()
	fmt.Println(text)
	ShowProgressBar(theme)
	time.Sleep(ProgressBarRefreshIntervalMs * time.Millisecond)
}
