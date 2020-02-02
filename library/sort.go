package library

// Methods helping to sort tracks in the library.

func (l *Library) Sort() {
	l.sortByAlpha()
}

func (l *Library) sortByAlpha() {

}

// func (l Library) Len() int           { return len(l.tracks) }
// func (l Library) Swap(i, j int)      { l.tracks[i],l.tracks[j]=l.tracks[j],l.tracks[i] }
// func (l Library) Less(i, j int) bool { return l.tracks[i].Title() < l.tracks[j] }
