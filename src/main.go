package playlist

func main() {

	p := InitPlaylist()
	p.AddSong("song1", 40)
	p.AddSong("song2", 40)
	p.AddSong("song3", 40)
	infoLog, f := initLogger()
	defer f.Close()
	infoLog.Println("INIT")
	managePlaylist(p, infoLog)
}
