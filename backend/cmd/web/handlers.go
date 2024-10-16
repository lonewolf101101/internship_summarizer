package main

import (
	"encoding/json"
	"io"
	"net/http"

	"undrakh.net/summarizer/pkg/common/oapi"
)

type Content struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.Copy(w, r.Body)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func getClientInput(w http.ResponseWriter, r *http.Request) {
	var input Content

	output := summarize(input)
	jsonResponse, err := json.Marshal(output)
	if err != nil {
		oapi.ClientError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
