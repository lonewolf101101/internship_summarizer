package main

import (
	"encoding/json"
	"net/http"
)

type Content struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (app *application) getClientInput(w http.ResponseWriter, r *http.Request) {
	var input Content

	output := app.summarize(input)
	jsonResponse, err := json.Marshal(output)
	if err != nil {
		app.clientError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
