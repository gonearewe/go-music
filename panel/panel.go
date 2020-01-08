package panel

import "fmt"

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
	eraseScreen()
	fmt.Println(LOGO)
}

func ShowCover(trackInfo string,theme ColorTheme, done chan bool ){
	eraseScreen()
	fmt.Println(RenderText(trackInfo, theme))	

	for{
		select{
		case <-done:
			eraseScreen()
			return
		default:
			break
		}
		
		ShowProgressBar(theme)
	}
}