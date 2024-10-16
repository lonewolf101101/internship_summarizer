package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Use(logRequest)
	r.Route("/", func(r chi.Router) {
		r.Get("/", home)
		r.Post("/summerize", getClientInput)
	})
	return r
}
