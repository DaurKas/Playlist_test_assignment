package main

import "container/list"

// SIGNALS
const PLAY = 1
const PAUSE = 2
const NEXT = 3
const PREV = 4

type Song struct {
	id       int
	name     string
	duration int
}

type Playlist struct {
	songs       *list.List
	hasStarted  bool
	currentSong Song
	ch          chan int
}

func songPlayer(p *Playlist) {
	// for ALL SONGS in playlist?
	n := 100 * 100 * 100 * p.currentSong.duration
	ch := p.ch
	for i := 0; i < n; i++ {
		select {
		case cmd := <-ch:
			if cmd == PAUSE {
				newCmd := <-ch
				if newCmd == PLAY {
					continue
				}
			} else if cmd == NEXT {
			} else if cmd == PREV {

			}
		default:
			continue
		}
	}
}

func (p *Playlist) AddSong(song Song) *Playlist {
	p.songs.PushBack(song)
	return p
}

func (p *Playlist) Play() {

	if !p.hasStarted {
		controlChannel := make(chan int)
		p.ch = controlChannel
		go songPlayer(p)
	} else {

	}

}

func (p *Playlist) Pause() {

}

func (p *Playlist) Next() {

}

func (p *Playlist) Prev() {

}
