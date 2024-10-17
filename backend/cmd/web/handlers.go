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

func summarizer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		oapi.CustomError(w, http.StatusBadRequest, "Unable to read body")
		return
	}
	defer r.Body.Close()

	var content Content
	err = json.Unmarshal(body, &content)
	if err != nil {
		oapi.CustomError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	resp, err := summarizeAPI(content)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	if resp.ErrMessage != "" {
		oapi.CustomError(w, resp.Code, resp.ErrMessage)
		return
	}

	// data := resp.Data
	// content.Summary := data.prompt

	response, err := json.Marshal(content)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
