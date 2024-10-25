package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"undrakh.net/summarizer/cmd/web/app"
)

// App holds the OAuth2 configurations

func routes() http.Handler {
	r := chi.NewRouter()
	r.Use(logRequest)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.With(authenticate).Route("/pub", func(r chi.Router) {
		r.Get("/logout", clearSession)

		r.Route("/auth", func(r chi.Router) {
			r.Route("/google", func(r chi.Router) {
				r.Get("/login", oauthLogin(app.GoogleOAuth2))
				r.Get("/callback", oauthCallback(app.GoogleOAuth2))
			})
		})

	})
	r.Route("/", func(r chi.Router) {
		r.Get("/", home)
		r.Post("/summarize", summarizer)
	})
	r.With(authenticate, requireAuth).Route("/api", func(r chi.Router) {
		r.Get("/me", me)
		r.Post("/me", updateUserInfo)
		r.HandleFunc("/ws", app.CustomerWSConnections.Handler)
		r.Get("/logout", logout)

		// TODO: Future plan: Admin
		r.With(requireAdmin).Route("/users", func(r chi.Router) {
			r.Get("/", getUsers)
			r.Post("/", editUser)
			r.With(setChosenUser).Route("/{UserID}", func(r chi.Router) {
				r.Get("/", getUser)
				r.Put("/", editUser)
				r.Delete("/", deleteUser)
			})
		})
	})

	return r
}
