package player

// ShowCover prints cover(info and progressbar) to the screen blockingly,
// thus this method requires started with a goroutine and closing
// param done(chan) directs its termination.
// func (p *Player) ShowCover(done <-chan struct{}) {
// 	panel.ShowCover(p.currentTrackInfo(), panel.RandomColorTheme(), done)
// 	return
// }

// func (p *Player) currentTrackInfo() string {
// 	track := *p.status.current
// 	return track.String()
// }
