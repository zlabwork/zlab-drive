package main

import (
	"context"
	"drive"
	"drive/app"
	"drive/app/api"
	"drive/app/middleware"
	"flag"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	// app.yaml
	bs, err := ioutil.ReadFile("../config/app.yaml")
	err = yaml.Unmarshal(bs, &drive.Cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// configs & router
	var wait time.Duration
	var dir string
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&dir, "dir", "assets", "the directory to serve files")
	flag.Parse()

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/", app.DefaultHandler)
	r.HandleFunc("/home", app.HomeHandler)
	r.HandleFunc("/path/{id:[0-9a-zA-Z_-]+}", api.PathHandler).Methods("GET")
	r.HandleFunc("/folders/{id:[0-9a-zA-Z_-]+}", app.DefaultHandler).Methods("GET")
	r.HandleFunc("/files/{id:[0-9a-zA-Z_-]+}", api.FilesHandler).Methods("GET")
	r.HandleFunc("/preview/{id:[0-9a-zA-Z_-]+}", app.PreviewHandler).Methods("GET")
	r.HandleFunc("/import", app.ImportHandler).Methods("GET")
	r.HandleFunc("/do/{id:[0-9a-zA-Z_-]+}", app.DefaultHandler).Methods("POST")
	r.HandleFunc("/holder/{id:[0-9x]+}", app.Placeholder).Methods("GET")
	r.HandleFunc("/upload", app.DefaultHandler).Methods("PUT")
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("../"+dir))))

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	log.Println("service is start")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
