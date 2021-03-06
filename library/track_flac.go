package library

import (
	"fmt"
)

type FLACTrack struct {
	baseFile file
	title    string // unlike filename, title comes from encoded tag info
	album    string
	artist   string // or artist list
	genre    string // or genre list
	year     string
	sum      string // sha256 sum of the track, serving as the unique id
}

// Title returns the title of a track.
func (f FLACTrack) Title() string {
	return f.title
}

// FileAddr returns complete address of the track file.
func (f FLACTrack) FileAddr() string {
	return f.baseFile.addr
}

// String wraps info of a track to a readable string.
func (f FLACTrack) String() string {
	s := fmt.Sprintln(f.title)
	s = fmt.Sprintln(s + f.artist)
	s = fmt.Sprintln(s + f.album + "   " + f.year)
	return fmt.Sprintln(s + f.genre)
}
