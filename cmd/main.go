package main

import (
	"fmt"
	"log/slog"
	"musiclibrary/db"
	"musiclibrary/internal/handlers"
	"musiclibrary/pkg/logger"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// @title Online Music Library API
// @version 1.0
// @description This is an API for managing a music library with CRUD operations and external API integration
// @host localhost:8080
// @BasePath /

func main() {
	logger.LogInit()
	slog.SetDefault(logger.Logger)
	db.DBInit()

	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	slog.Info(fmt.Sprintf("Server up with address: %v:%v", host, port))

	mux := http.DefaultServeMux

	mux.HandleFunc("GET /songs", handlers.GetSongHandler)
	mux.HandleFunc("GET /songs/{id}", handlers.GetDetailSongHandler)
	mux.HandleFunc("POST /songs", handlers.PostSongHandler)
	mux.HandleFunc("PUT /songs/{id}", handlers.UpdateSongHandler)
	mux.HandleFunc("DELETE /songs/{id}", handlers.DeleteSongHandler)

	mux.HandleFunc("/swagger/", serveSwagger)

	if err := http.ListenAndServe(host+":"+port, mux); err != nil {
		slog.Error(err.Error())
	}
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	if p == "" {
		p = "index.html"
	}
	p = filepath.Join("docs", p)
	http.ServeFile(w, r, p)
}
