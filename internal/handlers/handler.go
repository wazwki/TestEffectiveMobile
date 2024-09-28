package handlers

import (
	"TestEffectiveMobile/internal/models"
	"TestEffectiveMobile/internal/repository"
	"TestEffectiveMobile/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

// GetSongHandler godoc
// @Summary Get all songs
// @Description Get a list of all songs with optional filtering and pagination
// @Tags songs
// @Produce json
// @Param group query string false "Filter by group name"
// @Param song query string false "Filter by song name"
// @Param release_date query string false "Filter by release_date"
// @Param text query string false "Filter by song text"
// @Param limit query int false "Limit the number of results (default is 10)"
// @Param offset query int false "Offset the results for pagination (default is 0)"
// @Success 200 {array}	models.Song
// @Failure 500 "Status Internal Server Error 500"
// @Router /songs [get]
func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	release_date := r.URL.Query().Get("release_date")
	text := r.URL.Query().Get("text")

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	songs, err := repository.GetSong(group, song, release_date, text, limit, offset)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}

// GetDetailSongHandler godoc
// @Summary Get song details
// @Description Get detailed information about a song by its ID
// @Tags songs
// @Produce json
// @Param page query int false "Pagination"
// @Param id path int true "Song ID"
// @Success 200 {object} models.Song
// @Failure 400 "Status Bad Request 400"
// @Failure 500 "Status Internal Server Error 500"
// @Router /songs/{id} [get]
func GetDetailSongHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pageStr := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	var song *models.Song

	song, err = repository.GetDetailSong(id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	songObj := *song
	textsong := strings.Split(songObj.Text, "\n\n")

	if page <= len(textsong) {
		songObj.Text = textsong[page-1]
	} else {
		songObj.Text = textsong[len(textsong)-1]
	}

	if err = json.NewEncoder(w).Encode(&songObj); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// PostSongHandler godoc
// @Summary Add new song
// @Description Add a new song with external API enrichment
// @Tags songs
// @Accept json
// @Produce json
// @Param song body SongPart true "Add song"
// @Success 201 {object} models.SongPart
// @Failure 400 "Status Bad Request 400"
// @Failure 500 "Status Internal Server Error 500"
// @Router /songs [post]
func PostSongHandler(w http.ResponseWriter, r *http.Request) {
	var song models.Song

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fullsong, err := service.GetSongDetailFromOtherAPI(song)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = repository.PostSong(*fullsong)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateSongHandler godoc
// @Summary Update song
// @Description Update song details by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body Song true "Update song"
// @Success 200 {object} models.Song
// @Failure 400 "Status Bad Request 400"
// @Failure 500 "Status Internal Server Error 500"
// @Router /songs/{id} [put]
func UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var song models.Song

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := repository.UpdateSong(id, song)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteSongHandler godoc
// @Summary Delete song
// @Description Delete song by ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 204
// @Failure 400 "Status Bad Request 400"
// @Router /songs/{id} [delete]
func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := repository.DeleteSong(id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
