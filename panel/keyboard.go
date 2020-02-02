package panel

import (
	"bufio"
	"fmt"
	"os"
	"time"

	. "github.com/gonearewe/go-music/request"
)

// Here lie user-interaction functions, they occupy the terminal blockingly
// and expect users' input, they may give out signals according to users' instructions.

// RequestNewLibrary occupys terminal and requests users to type a new
// library's path and name.
func RequestNewLibrary() (libraryName, libraryPath string) {
	// wrap it since it's required when software initializes without library config.
	return setLibrary()
}

// listenForKeyboard blockingly listens for instructions given by key pression
// and changes the state of player accordingly by signaling player.
func listenForKeyboard(instruction string, outport chan<- Request) {
	switch instruction {
	case "q":
		EraseScreen()
		os.Exit(0)
	case "n":
		outport <- NewRequestToPlayer(RequestNextTrack)
	case "p":
		outport <- NewRequestToPlayer(RequestPrevTrack)
	case "x":
		showOptions(outport)
	}

}

// showOptions blockingly occupys screen for user interaction.
// NOTICE: It may throw exceptions.
func showOptions(outport chan<- Request) {
	switch getOption() {
	case 1:
		changeMode(outport)
	case 2:
		if libraryName, libraryPath := setLibrary(); libraryName != "" {
			outport <- NewRequestToMain(RequestSetNewLibrary, []string{libraryName, libraryPath})
		}
	}
	return
}

// getOption works in pair with showOptions.
func getOption() int {
	EraseScreen()

	fmt.Print(`
	Options:
		1: set player mode
		2: set library
		3: quit menu
	Please type your choice.
	>`)

	var input int
	fmt.Scanln(&input)
	switch input {
	case 1, 2, 3:
		return input
	default:
		return getOption()
	}
}

func changeMode(outport chan<- Request) {
	EraseScreen()
	var mode = [3]RequestType{RequestRandomMode, RequestRepeatMode, RequestSequentialMode}

	fmt.Print(`
	Set player mode:
		1: random mode
		2: repeat mode
		3: sequential mode
		4: quit menu
	Please type your choice.
	>`)

	var input int
	fmt.Scanln(&input)
	switch input {
	case 1, 2, 3:
		outport <- NewRequestToPlayer(mode[input-1])
	case 4:
	default:
		changeMode(outport)
	}
}

func setLibrary() (libraryName, libraryPath string) {
	EraseScreen()
	fmt.Print(`
	Set library: Type path:
	Please type folder path to your library, type 'q' to quit.
	>`)

	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Text() == "q" {
		return "", ""
	} else {
		libraryPath = sc.Text()
	}

	EraseScreen()
	// time.Sleep(100*time.Millisecond)

	fmt.Print(`
	Set library: Type name:
	Please type a name for your library, type 'q' to quit.
	>`)

	sc.Scan()
	if sc.Text() == "q" {
		return "", ""
	} else {
		libraryName = sc.Text()
	}

	EraseScreen()
	fmt.Println("Require reboot to start with new library ...")
	time.Sleep(1 * time.Second)

	if libraryName != "" && libraryPath != "" {
		return
	} else {
		return "", ""
	}
}
