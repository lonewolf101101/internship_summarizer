package main

import (
	"flag"
	"log"
	"net/http"

	"undrakh.net/summarizer/cmd/web/app"
)

// application structure will handle main initilzation and access and uses of functions
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
	app.InfoLog.Printf("Starting server on %s", *addr)
	app.ErrorLog.Fatal(srv.ListenAndServe())
}
