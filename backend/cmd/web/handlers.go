package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"undrakh.net/summarizer/cmd/web/app"
	"undrakh.net/summarizer/pkg/common/oapi"
)

type Content struct {
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

func logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r, "auth_user_id")
	app.Session.Remove(r, "oauth2_provider_name")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func clearSession(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r, "auth_user_id")
	app.Session.Remove(r, "oauth2_provider_name")
	oapi.SendResp(w, "OK")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func summarizer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read body: %v", err)
		oapi.CustomError(w, http.StatusBadRequest, "Unable to read body")
		return
	}
	defer r.Body.Close()

	var content Content
	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Printf("Invalid JSON: %v", err)
		oapi.CustomError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	resp, chatResp, err := summarizeAPI(content)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	if resp.ErrMessage != "" {
		oapi.CustomError(w, resp.Code, resp.ErrMessage)
		return
	}

	if len(chatResp.Choices) == 0 {
		oapi.CustomError(w, http.StatusInternalServerError, "No choices returned from summarization")
		return
	}

	content.Summary = chatResp.Choices[0].Message.Content
	response, err := json.Marshal(content)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
