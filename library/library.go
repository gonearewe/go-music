package library

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Library of tracks is based on a folder and contains tracks in his directory or subdirectories.
type Library struct {
	name   string // NOTICE: it's a unique id, it's determined by users or default, not the basename of the dirpath
	path   string // complete path
	tracks []Track
}

// NewLibrary initializes and returns a library with basic info and empty track list.
func NewLibrary(name, path string) (*Library, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	return &Library{
		name: name,
		path: filepath.Clean(path), // accept clean path string
	}, nil
}

func (l *Library) Scan() error {
	if files, err := ioutil.ReadDir(l.path); err != nil {
		return err
	} else {
		l.tracks = make([]Track, len(files)) // it's better to allocte first
	}

	filepath.Walk(l.path, func(path string, info os.FileInfo, err error) error {
		track, err := ParseTrack(path, info)
		if err != nil {
			return filepath.SkipDir // any other errors will terminate Walk function
		}
		panic(info.Name())
		l.tracks = append(l.tracks, track)
		return nil
	})

	if len(l.tracks) == 0 {
		return errors.New("library is empty")
	}

	fmt.Println(l.path, l.tracks)
	return nil
}

// func walk(root string, tracks []Track) {
// 	info, err := os.Lstat(root)
// 	if err != nil {
// 		err = walkFn(root, nil, err)
// 	} else {
// 		err = walk(root, info, walkFn)
// 	}
// }
