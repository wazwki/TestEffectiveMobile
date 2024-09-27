// @title Online Music Library API
// @version 1.0
// @description This is an API for managing a music library with CRUD operations and external API integration
// @host localhost:8080
// @BasePath /

package main

import (
	"TestEffectiveMobile/db"
	"TestEffectiveMobile/internal/handlers"
	"TestEffectiveMobile/pkg/logger"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	logger.LogInit()
	slog.SetDefault(logger.Logger)

	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := db.DBInit(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
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
