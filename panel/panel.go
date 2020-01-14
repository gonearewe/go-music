package panel

import (
	"fmt"
	"time"
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

func ShowLOGO() {
	EraseScreen()
	fmt.Println(LOGO)
}

func ShowCover(trackInfo string, theme ColorTheme, done <-chan struct{}) {
	const ProgressBarRefreshIntervalMs = 10 // how much time(ms) progressbar is refreshed

	var text=RenderText(trackInfo, theme)
	
	for {
		select {
		case <-done:
			EraseScreen()
			return
		default:
			break
		}

		EraseScreen()
		fmt.Println(text)
		ShowProgressBar(theme)
		time.Sleep(ProgressBarRefreshIntervalMs * time.Millisecond)
	}
}
