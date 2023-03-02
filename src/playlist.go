package main

import "container/list"

// SIGNALS
const PLAY = 1
const PAUSE = 2
const NEXT = 3
const PREV = 4

type Song struct {
	name     string
	duration int
}

type Playlist struct {
	songs       *list.List
	hasStarted  bool
	currentSong *list.Element
	ch          chan int
}

func initPlaylist() *Playlist {
	songs := list.New()
	ch := make(chan int)
	newPlaylist := Playlist{songs, false, nil, ch}
	return &newPlaylist
}

func songPlayer(p *Playlist) {
	for {
		var curSong Song
		if p.currentSong == nil {
			return
		}
		curSong = p.currentSong.Value.(Song)
		//imitating playing song
		n := 100 * 100 * 100 * curSong.duration
		ch := p.ch

		for i := 0; i < n; i++ {
			select {
			case cmd := <-ch:
				if cmd == PAUSE {
					select {
					case newCmd := <-ch:
						if newCmd == PLAY {
							continue
						} else if cmd == NEXT {
							p.currentSong = p.currentSong.Next()
							break
						} else if cmd == PREV {
							p.currentSong = p.currentSong.Prev()
							break

						}
					}
				} else if cmd == NEXT {
					p.currentSong = p.currentSong.Next()
					break
				} else if cmd == PREV {
					p.currentSong = p.currentSong.Prev()
					break

				}
			default:
				continue
			}
		}
	}
}

func (p *Playlist) AddSong(newName string, newDuration int) *Playlist {
	newSong := Song{newName, newDuration}
	p.songs.PushBack(newSong)
	return p
}

func (p *Playlist) Play() {

	if !p.hasStarted {
		go songPlayer(p)
	} else {
		p.ch <- PLAY
	}

}

func (p *Playlist) Pause() {
	p.ch <- PAUSE
}

func (p *Playlist) Next() {
	p.ch <- NEXT
}

func (p *Playlist) Prev() {
	p.ch <- PREV
}
