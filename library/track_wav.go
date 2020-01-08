package library

import (
	"os"
	"strings"
)

type WAVTrack struct {
	baseFile file
	title    string // for WAV track, title follows basefile's name
	sum      string // sha256 sum of the track, serving as the unique id
	// WAV format file doesn't necessarily contain tag information
}

// Title returns the title of a track.
func (w WAVTrack) Title() string {
	return w.title
}

// FileAddr returns complete address of the track file.
func (f WAVTrack) FileAddr() string {
	return f.baseFile.addr
}

// isWAVTrack checks the suffix of a filename and tells if it's a possible WAV file.
func isWAVTrack(fi os.FileInfo) bool {
	if strings.HasSuffix(fi.Name(), ".wav") {
		return true
	}

	return false
}

// String wraps info of a track to a readable string.
func (f WAVTrack) String() string {
	return f.title + "\n"
}
