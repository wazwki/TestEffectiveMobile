package handlers

import (
	"encoding/json"
	"log/slog"
	"musiclibrary/internal/models"
	"musiclibrary/internal/repository"
	"musiclibrary/internal/service"
	"net/http"
)

// GetSongHandler godoc
// @Summary Get all songs
// @Description Get a list of all songs with optional filtering and pagination
// @Tags songs
// @Produce json
// @Success 200 {array} Song
// @Failure 500 {object} ErrorResponse
// @Router /songs [get]
func GetSongHandler(w http.ResponseWriter, r *http.Request) {

	//вывод всей библиотеки(только поля group и song) с пагинацией и фильтрами
	w.Header().Set("Content-Type", "application/json")

}

// GetDetailSongHandler godoc
// @Summary Get song details
// @Description Get detailed information about a song by its ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} Song
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs/{id} [get]
func GetDetailSongHandler(w http.ResponseWriter, r *http.Request) {

	//вывод всех полей конкретного объекта
	w.Header().Set("Content-Type", "application/json")
}

// PostSongHandler godoc
// @Summary Add new song
// @Description Add a new song with external API enrichment
// @Tags songs
// @Accept json
// @Produce json
// @Param song body Song true "Add song"
// @Success 201 {object} Song
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
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
// @Success 200 {object} Song
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs/{id} [put]
func UpdateSongHandler(w http.ResponseWriter, r *http.Request) {

	//обновление объекта по id
}

// DeleteSongHandler godoc
// @Summary Delete song
// @Description Delete song by ID
// @Tags songs
// @Produce json
// @Param id path int true "Song ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs/{id} [delete]
func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {

	//удаление объекта по id
}
