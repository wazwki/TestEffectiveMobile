package service

import (
	"TestEffectiveMobile/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func GetSongDetailFromOtherAPI(song models.Song) (*models.Song, error) {
	var Details models.SongDetail
	apiURL := os.Getenv("API_URL")

	group := url.QueryEscape(song.GroupName)
	songName := url.QueryEscape(song.SongName)

	requestURL := fmt.Sprintf("%v/info?group=%v&song=%v", apiURL, group, songName)

	response, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("error making request to external API: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API returned non-200 status code: %v", response.StatusCode)
	}

	if err := json.NewDecoder(response.Body).Decode(&Details); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	song.ReleaseDate = Details.ReleaseDate
	song.Text = Details.Text
	song.Link = Details.Link

	return &song, nil
}
