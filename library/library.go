package library

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Library of tracks is based on a folder and contains tracks in his directory or subdirectories.
type Library struct {
	name   string // NOTICE: it's a unique id, it's determined by users or default, not the basename of the dirpath
	path   string // complete path
	tracks []Track
}

// NewLibrary initializes and returns a library with basic info and empty track list.
func NewLibrary(name, path string) (*Library, error) {
	if strings.Trim(name, " ") == "" {
		return nil, errors.New("name of spaces inacceptable")
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
//
// DEPRECATED: ScanWithRoutines proves to be faster than this one.
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

	return nil
}

func (l *Library) ScanWithRoutines() error {
	if files, err := ioutil.ReadDir(l.path); err != nil {
		return err
	} else {
		l.tracks = make([]Track, 0, len(files)) // it's better to allocte first
	}

	var wg = new(sync.WaitGroup)
	var track = make(chan Track)
	// every routine needs to acquire a token to start to work,
	// the length of tokens limits that the same number of routine can work.
	var tokens = make(chan struct{}, 20)
	wg.Add(1) // wait for mission dispatching
	walkWithRoutines(l.path, track, wg, tokens)
	wg.Done()

	// closer
	go func() {
		wg.Wait()
		close(track)
	}()

	// for range all tracks sent by working routines
	for t := range track {
		l.tracks = append(l.tracks, t)
	}

	if len(l.tracks) == 0 {
		return errors.New("library is empty")
	}
	return nil
}

func (l *Library) GetTrackByID(id int) *Track {
	if id > len(l.tracks)-1 {
		panic("track id out of range") // let it crash
	}

	return &l.tracks[id]
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

func (l *Library) NumTracks() int {
	if l.name == "" {
		panic("uninitialized library")
	}
	return len(l.tracks)
}

// walk searches a folder and its sub-folder recursively for tracks.
func walk(path string, tracks *[]Track) {
	entries, ok := dirEntries(path)
	if !ok {
		return
	}

	for _, e := range entries {
		subpath := filepath.Join(path, e.Name())

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

// walkWithRoutines trys to traverse a dictionary recursively and
// parses every track by starting a routine to handle it, though
// number of routines is limited through tokens for scheduling costs time.
func walkWithRoutines(path string, track chan<- Track, wg *sync.WaitGroup, tokens chan struct{}) {
	file, err := os.Stat(path)
	if err != nil {
		return
	}

	if !file.IsDir() {
		wg.Add(1)
		go func(file os.FileInfo) {
			defer wg.Done()
			defer func() { <-tokens }()
			tokens <- struct{}{}

			if t, err := ParseTrack(path, file); err == nil {
				track <- t
			}
		}(file)
	} else {
		subfiles, err := ioutil.ReadDir(path)
		if err != nil {
			return
		}

		for _, f := range subfiles {
			subpath := filepath.Join(path, f.Name())
			walkWithRoutines(subpath, track, wg, tokens)
		}
	}
}
