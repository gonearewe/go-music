package player

import (
	"fmt"
	"bufio"
	"os"

	"github.com/gonearewe/go-music/panel"
)

// ListenForKeyboard blockingly listens for instructions given by key pression
// and changes the state of player accordingly.
func (p *Player) ListenForKeyboard() {
	var input = bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanRunes)
	for {
		if !input.Scan() {
			continue
		}
		instruction := input.Text()
		fmt.Println(instruction)
		switch instruction {
		case "q":
			panel.EraseScreen()
			os.Exit(0)
		case "n":
			p.PlayNextTrack()
		case "p":
			p.PlayPrevTrack()
		case "x":
			p.ShowOptions()
		}
	}
}

func (p *Player)ShowOptions(){
	if p.done!=nil{
		close(p.done)
	}else{
		panel.EraseScreen()
	}

	switch getOption(){
	case 1:
		p.changeMode()
	case 2:
		p.setLibrary()
	case 3:
		return
	}
}

func getOption() int{
	panel.EraseScreen()

	fmt.Println(`
	Options:
		1: set player mode
		2: set library
		3: quit menu
	Please type your choice.
	>`)

	var input int
	fmt.Scanln(&input)
	switch input{
	case 1,2,3:
		return input
	default:
		return getOption()
	}
}

func (p *Player)changeMode(){
	panel.EraseScreen()
	var mode=[3]PlayerMode{RandomMode,RepeatMode,SequentialMode}

	fmt.Println(`
	Set player mode:
		1: random mode
		2: repeat mode
		3: sequential mode
		4: quit menu
	Please type your choice.
	>`)

	var input int
	fmt.Scanln(&input)
	switch input{
	case 1,2,3:
		p.SetMode(mode[input])
	default:
		p.changeMode()
	}
}
