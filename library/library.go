package library

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Library of tracks is based on a folder and contains tracks in his directory or subdirectories.
type Library struct {
	name   string // NOTICE: it's a unique id, it's determined by users or default, not the basename of the dirpath
	path   string // complete path
	tracks []Track
}

// NewLibrary initializes and returns a library with basic info and empty track list.
func NewLibrary(name, path string) (*Library, error) {
	if strings.Trim(name, " ")==""{
		return nil,errors.New("name of spaces inacceptable")
	}

	path = filepath.Clean(path)
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	return &Library{
		name:   name,
		path:   path, // accept clean path string
		tracks: []Track{},
	}, nil
}

// Scan scans tracks for a initialized library(with a path referring to a music folder), 
// scanning a empty library will results in a error.
func (l *Library) Scan() error {
	if files, err := ioutil.ReadDir(l.path); err != nil {
		return err
	} else {
		l.tracks = make([]Track, 0, len(files)) // it's better to allocte first
	}

	walk(l.path, &l.tracks)

	if len(l.tracks) == 0 {
		return errors.New("library is empty")
	}

	fmt.Println(l.path, l.tracks)
	return nil
}

// String wraps info of a library to a readable string.
func (l *Library) String() string {
	str := "name: " + l.name + "\n" + "path: " + l.path + "\n"
	// tracksStr := make([]string, len(l.tracks))
	var tracksStr []string
	for _, track := range l.tracks {
		tracksStr = append(tracksStr, track.Title())
	}
	return str + "tracks: \n" + strings.Join(tracksStr, "\n")
}

func (l *Library)NumTracks()int{
	if l.name==""{
		panic("uninitialized library")
	}
	return len(l.tracks)
}

// TODO: improve with goroutine supports.
// walk searchs a folder and its sub-folder recursively for tracks.
func walk(path string, tracks *[]Track) {
	entries, ok := dirEntries(path)
	if !ok {
		return
	}

	for _, e := range entries {
		subpath:=filepath.Join(path, e.Name())

		if e.IsDir() {
			walk(subpath, tracks)
			return
		}

		if track, err := ParseTrack(subpath, e); err == nil {
			*tracks = append(*tracks, track)
		}
	}

	return
}

// dirEntries generates a slice of direct contents of given folder,
// and if given path is not a folder, returns false.
func dirEntries(path string) ([]os.FileInfo, bool) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, false
	}

	return entries, true
}
