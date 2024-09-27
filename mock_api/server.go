package mock_api

import (
	"TestEffectiveMobile/internal/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func MockApiInit() error {
	host := os.Getenv("API_URL")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8081"
	}

	slog.Info(fmt.Sprintf("Mock API up with adress: %v:%v", host[7:], port))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /info", func(w http.ResponseWriter, r *http.Request) {
		text := "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
		link := "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
		release_date := "2006-07-16"
		details := models.SongDetail{ReleaseDate: release_date, Text: text, Link: link}
		json.NewEncoder(w).Encode(&details)
		w.Header().Set("Content-Type", "application/json")
	})

	if err := http.ListenAndServe(host[7:]+":"+port, mux); err != nil {
		return err
	}
	return nil
}
