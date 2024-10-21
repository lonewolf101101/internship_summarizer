package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Use(logRequest)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Get("/", home)
		r.Post("/summarize", summarizer)
	})
	return r
}
