package panel

import (
	"math/rand"
	"strings"

	. "github.com/fatih/color"
)

var (
	musicSigns  = [4]string{"\u266A", "\u266B", "\u266C", "\u2669"}
	progressBar ProgressBar
)

type ProgressBar struct {
	bar             string // bar is a string consisting of musicSigns
	isGettingLonger bool   // the way progress bar goes, gets longer or shorter?
	progress        int    // ranging from 0 to len(bar)-1, records current progress
}

func ShowProgressBar(theme ColorTheme) {
	EraseCurrentLine()
	// at the start
	if progressBar.progress == 0 && !progressBar.isGettingLonger {
		len := rand.Intn(20) + 10 // NOTE: 10 <= len <= 30

		s := make([]string, len)
		for i := range s {
			s[i] = musicSigns[rand.Intn(4)]
		}
		progressBar.bar = strings.Join(s, "")

		progressBar.isGettingLonger = true
	}

	// at the end
	if progressBar.progress == len(progressBar.bar)-1 && progressBar.isGettingLonger {
		progressBar.isGettingLonger = false
	}

	// prints bar to the screen
	if progressBar.isGettingLonger {
		progressBar.progress++
		New(theme[0]).Print(progressBar.bar[:progressBar.progress+1])
	} else {
		progressBar.progress--
		New(theme[1]).Print(progressBar.bar[:progressBar.progress+1])
	}
}
