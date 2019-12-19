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

func (l *Library) Scan() error {
	if files, err := ioutil.ReadDir(l.path); err != nil {
		return err
	} else {
		l.tracks = make([]Track, len(files)) // it's better to allocte first
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

func walk(path string, tracks *[]Track) {
	entries, ok := dirEntries(path)
	if !ok {
		return
	}

	for _,e:=range entries{
		if e.IsDir(){
			walk(filepath.Join(path,e.Name()), tracks)
			return
		}

		if track, err := ParseTrack(path, e);err == nil {
			*tracks = append(*tracks, track)
		}
	}

	return
}

func dirEntries(path string) ([]os.FileInfo, bool) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, false
	}

	return entries, true
}
