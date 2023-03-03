package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
}

func (app *application) play(w http.ResponseWriter, r *http.Request) {
	app.playlist.Play(app.infoLog)
}
func (app *application) addSong(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		duration, _ := strconv.Atoi(r.FormValue("duration"))
		app.playlist.AddSong(name, duration)

		filename, err := os.Open("data.json")
		if err != nil {
			app.errorLog.Fatal(err)
		}

		defer filename.Close()

		jsonData, jsonErr := json.Marshal(app.playlist)

		if jsonErr != nil {
			app.errorLog.Fatal(jsonErr)
		}
		go ioutil.WriteFile("data.json", jsonData, 0644)

	}
}
func (app *application) deleteSong(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		duration, _ := strconv.Atoi(r.FormValue("duration"))
		app.playlist.DeleteSong(name, duration)
	}
}
func (app *application) next(w http.ResponseWriter, r *http.Request) {
	app.playlist.Next()
}
func (app *application) prev(w http.ResponseWriter, r *http.Request) {
	app.playlist.Prev()
}
func (app *application) pause(w http.ResponseWriter, r *http.Request) {
	app.playlist.Pause()
}
