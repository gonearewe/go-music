package panel

import (
	"fmt"
	"os"

	. "github.com/gonearewe/go-music/request"
)

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
		setLibrary()
	case 3:
		return
	}
}

// SetLibrary is required when software initializes without library config.
// NOTICE: It may throw exceptions.
func SetLibrary() {
	setLibrary()
}

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

func setLibrary() {
	EraseScreen()
	fmt.Println("TODO")
	// fmt.Println(`
	// Set library: Type path:
	// Please type folder path to your library, type 'q' to quit.
	// >`)

	// var input, libraryPath, libraryName string
	// fmt.Scanln(&input)
	// if input == "q" {
	// 	return
	// } else {
	// 	libraryPath = input
	// }

	// EraseScreen()
	// fmt.Println(`
	// Set library: Type name:
	// Please type a name for your library, type 'q' to quit.
	// >`)

	// fmt.Scanln(&input)
	// if input == "q" {
	// 	return
	// } else {
	// 	libraryName = input
	// }

	// EraseScreen()
	// fmt.Println("Loading Library... Please wait... ")

	// lib, err := library.NewLibrary(libraryName, libraryPath)
	// if err != nil {
	// 	panic(err)
	// }
	// err = lib.Scan()
	// if err != nil {
	// 	panic(err)
	// }

	// p = NewPlayer(lib)
}
