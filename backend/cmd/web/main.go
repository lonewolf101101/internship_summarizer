package main

import (
	"flag"
	"net/http"

	"undrakh.net/summarizer/cmd/web/app"
)

// application structure will handle main initilzation and access and uses of functions
func main() {
	addr := flag.String("addr", ":3300", "HTTP network address")
	// secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()
	app.Init()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLog,
		Handler:  routes(),
	}
	app.InfoLog.Printf("Starting server on %s", *addr)
	app.ErrorLog.Fatal(srv.ListenAndServe())
}
