/*
* playlist acts like a limited-size stack, when you push new one and it's full,
* the oldest one will be removed, so it can serve as a cache storing recently played 
* tracks so you can access previous track.
*/

package player

import (
	"errors"
	"container/list"

	"github.com/gonearewe/go-music/library"
)

const PLAYLIST_MAX_SIZE=10

type playList struct {
	*list.List
}

type playListElem struct {
	track library.Track
	id    int // ID of this track in Library(assuming it won't change)
}

func newPlayList() *playList {
	return &playList{list.New()}
}

// push pushs a track and its id into the playlist.
func (l *playList) push(track library.Track, id int) {
	l.PushFront(playListElem{track, id})
	if l.Len()>PLAYLIST_MAX_SIZE{
		l.Remove(l.Back())
	}
}

// peeks returns the track on the top(current one) and reports error when it's empty.
func (l *playList) peek()(track library.Track,id int,err error){
	cur:=l.Front()
	if cur==nil{
		return nil,0,errors.New("empty playlist")
	}

	if e,ok:=cur.Value.(playListElem);ok{
		return e.track,e.id, nil
	}else {
		return nil,0,errors.New("cast to playListElem: did you put something else into the list?")
	}
}

// pop pops a track out and reports error when it's empty.
func (l *playList) pop()(library.Track,error){
	if e,_,err:=l.peek();err!=nil{
		return nil,err
	}else {
		l.Remove(l.Front()) // pop out
		return e,nil
	}
}

