package player

import (
	"github.com/gonearewe/go-music/panel"
)

func (p *Player) ShowCover(theme panel.ColorTheme, done chan bool) {
	panel.ShowCover(p.currentTrackInfo(), theme, done)
	return
}

func (p *Player) currentTrackInfo() string {
	track := *p.status.current
	return track.String()
}
