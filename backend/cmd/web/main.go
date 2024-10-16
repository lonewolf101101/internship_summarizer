package main

import (
	"flag"
	"log"
	"net/http"

	"undrakh.net/summarizer/cmd/web/app"
)

// application structure will handle main initilzation and access and uses of functions
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":3300", "HTTP network address")
	flag.Parse()
	app.Init()
	log.Println("Starting server on :3300")

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.ErrorLog,
		Handler:  routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
