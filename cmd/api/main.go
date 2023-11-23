package main

import (
	"flag"

	"net/http"
	"time"

	"log"
	"os"
)

type App struct {
	DB   string
	PORT string
	log  *log.Logger
}

func main() {

	logger := log.New(os.Stdout, "", log.Ltime|log.Ldate)
	app := App{log: logger}

	flag.StringVar(&app.DB, "d", "local", "database string")
	flag.StringVar(&app.PORT, "port", ":5000", "port")
	flag.Parse()

	srv := &http.Server{
		Addr:         app.PORT,
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.log.Printf("starting server on %s", app.PORT)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
