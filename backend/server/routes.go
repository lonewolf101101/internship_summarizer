package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standartMiddleware := alice.New(app.logRequest)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Post("/summerize", http.HandlerFunc(app.getClientInput))
	return standartMiddleware.Then(mux)
}
