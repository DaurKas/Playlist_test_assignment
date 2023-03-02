package main

import (
	"fmt"
	"log"
	"os"
)

func initLogger() *log.Logger {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	return infoLog
}

func managePlaylist(p *Playlist, infoLog *log.Logger) {
	var cmd string
	for {
		fmt.Scan(&cmd)
		switch cmd {
		case "PLAY":
			p.Play()
		case "PAUSE":
			p.Pause()
		case "NEXT":
			p.Next()
		case "PREV":
			p.Prev()
		case "ADD":
			fmt.Println("PRINT NEW SONG INFO IN FORMAT: %NAME %DURATION")
			var newName string
			var newDuration int
			fmt.Scan(&newName, &newDuration)
			p = p.AddSong(newName, newDuration)
		}

	}
}
