package playlist

import (
	"fmt"
	"log"
	"os"
)

func initLogger() (*log.Logger, *os.File) {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	return infoLog, f
}

func managePlaylist(p *Playlist, infoLog *log.Logger) {
	var cmd string
	for {
		fmt.Scan(&cmd)
		switch cmd {
		case "PLAY":
			infoLog.Print(p.Play(infoLog))
		case "PAUSE":
			infoLog.Println(p.Pause())
		case "NEXT":
			infoLog.Println(p.Next())
		case "PREV":
			infoLog.Println(p.Prev())
		case "ADD":
			fmt.Println("PRINT NEW SONG INFO IN FORMAT: %NAME %DURATION")
			var newName string
			var newDuration int
			fmt.Scan(&newName, &newDuration)
			p = p.AddSong(newName, newDuration)
		case "EXIT":
			return
		}

	}
}
