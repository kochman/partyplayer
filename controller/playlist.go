package controller

type Playlist struct {
	Name  string
	Songs []Song
}

type Song struct {
}

func (p *Playlist) AddSong(song Song) {
	p.Songs = append(p.Songs, song)
}
