package playlist

import (
	"container/list"
	"log"
)

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
	hasEnded    bool
	currentSong *list.Element
	ch          chan int
}

func InitPlaylist() *Playlist {
	songs := list.New()
	ch := make(chan int)
	newPlaylist := Playlist{songs, false, false, nil, ch}
	return &newPlaylist
}

func (p *Playlist) AddSong(newName string, newDuration int) *Playlist {
	newSong := Song{newName, newDuration}
	p.songs.PushBack(newSong)
	if p.songs.Len() == 1 {
		p.currentSong = p.songs.Front()
	}
	return p
}

func (p *Playlist) Play(infoLog *log.Logger) string {

	if !p.hasStarted || p.hasEnded {
		p.hasEnded = false
		p.hasStarted = true
		go songPlayer(p, infoLog)
	} else {

		p.ch <- PLAY
	}
	return "PLAY"

}

func (p *Playlist) Pause() string {
	if p.hasEnded || !p.hasStarted {
		return "Playlist is not playing"
	}
	p.ch <- PAUSE
	return "PAUSED"
}

func (p *Playlist) Next() string {
	if p.hasEnded || !p.hasStarted {
		return "Playlist is not playing"
	}
	p.ch <- NEXT
	return "NEXT"
}

func (p *Playlist) Prev() string {
	if p.hasEnded || !p.hasStarted {
		return "Playlist is not playing"
	}
	p.ch <- PREV
	return "PREV"
}

func (p *Playlist) DeleteSong(newName string, newDuration int) *Playlist {
	newSong := Song{newName, newDuration}
	p.songs.PushBack(newSong)
	if p.songs.Len() == 1 {
		p.currentSong = p.songs.Front()
	}
	return p
}

func songPlayer(p *Playlist, infoLog *log.Logger) {
	for {
		var curSong Song
		skipFlag := false
		pauseFlag := false
		if p.currentSong == nil {
			infoLog.Println("PLAYLIST HAS ENDED")
			p.hasEnded = true
			p.currentSong = p.songs.Front()
			return
		}
		curSong = p.currentSong.Value.(Song)
		n := 1000 * 100 * 1000 * curSong.duration
		ch := p.ch
		infoLog.Println("PLAYING SONG:", curSong.name)
		for i := 0; i < n; i++ {
			if pauseFlag {
				i--
			}

			select {
			case cmd := <-ch:
				if cmd == PAUSE {
					pauseFlag = true
				} else if cmd == NEXT {
					p.currentSong = p.currentSong.Next()
					pauseFlag = false
					skipFlag = true
					break
				} else if cmd == PREV {
					p.currentSong = p.currentSong.Prev()
					pauseFlag = false
					skipFlag = true
					break

				} else if cmd == PLAY {
					pauseFlag = false
				}
			default:
				continue
			}
			if skipFlag {
				break
			}
			infoLog.Println("AFTER SWITCH")
		}
		if !skipFlag {
			p.currentSong = p.currentSong.Next()
		}
		infoLog.Println("END OF SONG")
		skipFlag = false

	}
}
