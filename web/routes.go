package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/playlist/play", app.play)
	mux.HandleFunc("/playlist/pause", app.pause)
	mux.HandleFunc("/playlist/next", app.next)
	mux.HandleFunc("/playlist/prev", app.prev)
	mux.HandleFunc("/playlist/add", app.addSong)
	mux.HandleFunc("/playlist/remove", app.deleteSong)

	return mux
}
