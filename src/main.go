package main

func main() {

	p := initPlaylist()
	p.AddSong("song1", 40)
	p.AddSong("song2", 40)
	p.AddSong("song3", 40)
	infoLog := initLogger()
	managePlaylist(p, infoLog)
}
