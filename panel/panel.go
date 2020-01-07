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
