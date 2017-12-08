package controller

import (
	"sync"

	"github.com/kochman/partyplayer/config"
)

type Controller struct {
	m         sync.Mutex
	playlists map[string]*Playlist
}

func New(cfg *config.Config) (*Controller, error) {
	c := &Controller{}
	return c, nil
}

func (c *Controller) Playlist(name string) *Playlist {
	c.m.Lock()
	defer c.m.Unlock()
	if playlist, ok := c.playlists[name]; ok {
		return playlist
	}
	p := &Playlist{}
	c.playlists[name] = p
	return p
}

func (c *Controller) Playlists() []*Playlist {
	playlists := []*Playlist{}
	c.m.Lock()
	defer c.m.Unlock()
	for _, playlist := range c.playlists {
		playlists = append(playlists, playlist)
	}
	return playlists
}
