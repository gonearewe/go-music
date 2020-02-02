package library

import (
	"github.com/gonearewe/go-music/config"
)

//  ScanLibraryAndSaveConfig sets up a new library and scans it before saves track infos into config file.
func ScanLibraryAndSaveConfig(libName, libPath string) error {
	lib, err := NewLibrary(libName, libPath)
	if err != nil {
		return err
	}
	err = lib.ScanWithRoutines()
	if err != nil {
		return err
	}
	libconfig := NewLibraryConfiguration(lib)

	return config.SaveConfigInWorkDir(libconfig)
}

// NewLibraryConfiguration initializes a *config.LibraryConfiguration from given *Library.
func NewLibraryConfiguration(lib *Library) *config.LibraryConfiguration {
	libconfig := &config.LibraryConfiguration{
		Libraries: make([]config.Library, 1),
	}

	libconfig.Libraries[0].Name = lib.name
	libconfig.Libraries[0].Path = lib.path
	libconfig.Libraries[0].Tracks = make([]config.Track, len(lib.tracks))
	passTrackInfo(&libconfig.Libraries[0].Tracks, &lib.tracks)

	return libconfig
}

// NewLibraryFromConfig initializes a library from config file, which is faster than scanning local directory.
func NewLibraryFromConfig(libconfig *config.LibraryConfiguration) *Library {
	lib := &Library{
		name:   libconfig.Libraries[0].Name,
		path:   libconfig.Libraries[0].Path,
		tracks: make([]Track, len(libconfig.Libraries[0].Tracks)),
	}

	loadTrackInfo(&lib.tracks, &libconfig.Libraries[0].Tracks)

	return lib
}

// passTrackInfo passes track infos from []Track to []config.Track.
func passTrackInfo(cts *[]config.Track, tracks *[]Track) {
	for i, track := range *tracks {
		ct := &(*cts)[i]

		if t, ok := track.(FLACTrack); ok {
			ct.Format = "FLAC"
			ct.Album = t.album
			ct.Artist = t.artist
			ct.Genre = t.genre
			ct.Sum = t.sum
			ct.Title = t.title
			ct.Year = t.year

			ct.BaseFile.Addr = t.baseFile.addr
			ct.BaseFile.Name = t.baseFile.name
			ct.BaseFile.Size = t.baseFile.size
		} else if t, ok := track.(WAVTrack); ok {
			ct.Format = "WAV"
			ct.Title = t.title
			ct.Sum = t.sum

			ct.BaseFile.Addr = t.baseFile.addr
			ct.BaseFile.Name = t.baseFile.name
			ct.BaseFile.Size = t.baseFile.size
		} else {
			panic("unknown format")
		}

	}
}

// loadTrackInfo loads track infos from []config.Track to []Track.
func loadTrackInfo(tracks *[]Track, cts *[]config.Track) {
	for i, ct := range *cts {
		if ct.Format == "FLAC" {
			(*tracks)[i] = FLACTrack{
				album:  ct.Album,
				artist: ct.Artist,
				genre:  ct.Genre,
				sum:    ct.Sum,
				title:  ct.Title,
				year:   ct.Year,
				baseFile: file{
					addr: ct.BaseFile.Addr,
					name: ct.BaseFile.Name,
					size: ct.BaseFile.Size,
				},
			}
		} else if ct.Format == "WAV" {
			(*tracks)[i] = WAVTrack{
				sum:   ct.Sum,
				title: ct.Title,
				baseFile: file{
					addr: ct.BaseFile.Addr,
					name: ct.BaseFile.Name,
					size: ct.BaseFile.Size,
				},
			}
		} else {
			panic("unknown format")
		}

	}
}
