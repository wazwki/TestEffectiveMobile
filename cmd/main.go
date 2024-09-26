package main

import (
	"log"
	"musiclibrary/internal/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	LogInit()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load env", err)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.DefaultServeMux

	mux.HandleFunc("GET /info", handlers.GetSongHandler)
	mux.HandleFunc("GET /info/{id}", handlers.GetDetailSongHandler)
	mux.HandleFunc("POST /info", handlers.PostSongHandler)
	mux.HandleFunc("PUT /info/{id}", handlers.UpdateSongHandler)
	mux.HandleFunc("DELETE /info/{id}", handlers.DeleteSongHandler)

	if err = http.ListenAndServe(host+":"+port, mux); err != nil {
		log.Fatal("Server cant start", err)
	}
}
