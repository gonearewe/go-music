package library

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/jcjordyn130/gotag"
)

const LEAST_TRACK_FILE_SIZE = 1024 // least size(byte)
var ValidTrackFileSuffixes = [2]string{".wav", ".flac"}

type Track interface{}

type WAVTrack struct {
	baseFile file
	sum      string // sha256 sum of the track, serving as the unique id
	// WAV format file doesn't necessarily contain tag information
}

type FLACTrack struct {
	baseFile file
	title    string // unlike filename, title are
	album    string
	artist   string // or artist list
	genre    string // or genre list
	year     string
	sum      string // sha256 sum of the track, serving as the unique id
}

type file struct {
	addr string // complete address 
	name string // file basename
	size int64 // how many bytes the file is
}

func ParseTrack(path string, fi os.FileInfo) (Track, error) {
	if !isValidTrack(fi) {
		return nil, errors.New("invalid track file")
	}

	metadata, err := gotag.Open(path)
	if err != nil {
		if isWAVTrack(fi) {
			return WAVTrack{
				baseFile: file{
					addr: path,
					name: fi.Name(),
					size: fi.Size(),
				},
			}, nil
		}

		return nil, err
	}

	// we don't care about the err so if the field isn't found, it's simply empty
	title, _ := metadata.Title()
	album, _ := metadata.Album()
	artists, _ := metadata.Artist()
	genres, _ := metadata.Genre()
	year, _ := metadata.Year()
	sum,_:=metadata.Sum()

	// TODO: only supports flac and wav for now
	flactrack := FLACTrack{
		baseFile: file{
			addr: path,
			name: fi.Name(),
			size: fi.Size(),
		},
		title:  title,
		album:  album,
		artist: strings.Join(artists, ", "), // multible artists are separated by ", "
		genre:  strings.Join(genres, ", "),  // multible genres are separated by ", "
		year:   string(year),
		sum:sum,
	}

	// DEBUG
	fmt.Println(flactrack)
	return flactrack, nil
}

// sum returns the sha256 sum of a file.
func sum(fileaddr string) (string, error) {
	// Open the file.
	file, err := os.Open(fileaddr)
	if err != nil {
		return "", err
	}

	// Defer closing it.
	defer file.Close()

	// Make a new hash object.
	hash := sha256.New()

	// Copy the file to the hasher object.
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// Calculate the hash and return it.
	return hex.EncodeToString(hash.Sum(nil)), err
}

// isValidTrack tells if a file is a possible track by provided file info.
func isValidTrack(fi os.FileInfo) bool {
	if fi.IsDir() || fi.Size() < LEAST_TRACK_FILE_SIZE {
		return false
	}

	for _, suffix := range ValidTrackFileSuffixes {
		if strings.HasSuffix(fi.Name(), suffix) {
			return true
		}
	}

	return false
}

// isWAVTrack checks the suffix of a filename and tells if it's a possible WAV file.
func isWAVTrack(fi os.FileInfo) bool {
	if strings.HasSuffix(fi.Name(), ".wav") {
		return true
	}

	return false
}
