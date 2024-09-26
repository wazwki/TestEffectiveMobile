package repository

import "musiclibrary/internal/models"

func GetSong() ([]models.Song, error) {
	//
	return []models.Song{}, nil
}

func GetDetailSong(id int) (models.Song, error) {
	//
	return models.Song{}, nil
}

func PostSong(s models.Song) error {
	//
	return nil
}

func UpdateSong(id int) error {
	//
	return nil
}

func DeleteSong(id int) error {
	//
	return nil
}
