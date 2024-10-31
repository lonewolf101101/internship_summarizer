package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"undrakh.net/summarizer/cmd/web/app"
)

// application structure will handle main initilzation and access and uses of functions
func main() {
	addr := flag.String("addr", ":3300", "HTTP network address")
	// secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()
	app.Init()
	defer app.Close()

	closeOnSignalInterrupt(app.Close)

	// panicOnError(app.DB.AutoMigrate(
	// 	new(userman.User),
	// 	new(roleman.Role),
	// ))

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     app.ErrorLog,
		Handler:      routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}
	app.InfoLog.Printf("Starting server on %s", *addr)
	app.ErrorLog.Fatal(srv.ListenAndServe())
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func closeOnSignalInterrupt(cleanupFunc func()) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		cleanupFunc()
		os.Exit(0)
	}()
}
