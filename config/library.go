package config

import "github.com/pelletier/go-toml"

type LibraryConfiguration struct {
	Libraries []Library
}

type Library struct {
	Name   string // NOTICE: it's a unique id, it's determined by users or default, not the basename of the dirpath
	Path   string // complete path
	Tracks []Track
}

type Track struct {
	BaseFile file
	Format   string
	Title    string // unlike filename, title comes from encoded tag info
	Album    string
	Artist   string // or artist list
	Genre    string // or genre list
	Year     string
	Sum      string // sha256 sum of the track, serving as the unique id
}

type file struct {
	Addr string // complete address
	Name string // file basename
	Size int64  // how many bytes the file is
}

func (l *LibraryConfiguration) FileName() string {
	return "libraries.toml"
}

func (l *LibraryConfiguration) Marshal() ([]byte, error) {
	return toml.Marshal(*l)
}

func (l *LibraryConfiguration) Unmarshal(data []byte) error {
	return toml.Unmarshal(data, l)
}
