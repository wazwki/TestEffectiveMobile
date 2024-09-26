package service

import (
	"encoding/json"
	"fmt"
	"musiclibrary/internal/models"
	"net/http"
	"os"
)

func GetSongDetailFromOtherAPI(song models.Song) (*models.Song, error) {
	var Details models.SongDetail
	api_url := os.Getenv("API_URL")
	response, err := http.Get(fmt.Sprintf("%v?group=%v&song=%v", api_url, song.GroupName, song.SongName))
	if err != nil {
		return nil, err
	}
	json.NewDecoder(response.Body).Decode(&Details)

	song.ReleaseDate = Details.ReleaseDate
	song.Text = Details.Text
	song.Link = Details.Link

	return &song, nil
}
