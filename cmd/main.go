// @title Online Music Library API
// @version 1.0
// @description This is an API for managing a music library with CRUD operations and external API integration
// @host localhost:8080
// @BasePath /

package main

import (
	"TestEffectiveMobile/db"
	"TestEffectiveMobile/internal/handlers"
	"TestEffectiveMobile/mock_api"
	"TestEffectiveMobile/pkg/logger"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

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

	mux := http.DefaultServeMux

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", host, port),
		Handler: mux,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	mux.HandleFunc("GET /songs", handlers.GetSongHandler)
	mux.HandleFunc("GET /songs/{id}", handlers.GetDetailSongHandler)
	mux.HandleFunc("POST /songs", handlers.PostSongHandler)
	mux.HandleFunc("PUT /songs/{id}", handlers.UpdateSongHandler)
	mux.HandleFunc("DELETE /songs/{id}", handlers.DeleteSongHandler)

	mux.HandleFunc("/swagger/", serveSwagger)

	go func() {
		slog.Info(fmt.Sprintf("Server up with address: %v:%v", host, port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
		}
	}()

	go func() {
		if err := mock_api.MockApiInit(); err != nil {
			slog.Error(err.Error())
		}
	}()

	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	} else {
		slog.Info("Server gracefully stopped")
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
