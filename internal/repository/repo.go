package repository

import (
	"TestEffectiveMobile/db"
	"TestEffectiveMobile/internal/models"
	"strconv"
	"strings"
)

func GetSong(group, song, release_date, text string, limit, offset int) ([]models.Song, error) {
	songs := []models.Song{}

	query := `SELECT id, group_name, song_name, release_date, text, link FROM songs WHERE 1=1`
	var args []interface{}
	var conditions []string

	if group != "" {
		conditions = append(conditions, "group_name ILIKE $"+strconv.Itoa(len(args)+1))
		args = append(args, "%"+group+"%")
	}

	if song != "" {
		conditions = append(conditions, "song_name ILIKE $"+strconv.Itoa(len(args)+1))
		args = append(args, "%"+song+"%")
	}

	if release_date != "" {
		conditions = append(conditions, "release_date = $"+strconv.Itoa(len(args)+1))
		args = append(args, release_date)
	}

	if text != "" {
		conditions = append(conditions, "text ILIKE $"+strconv.Itoa(len(args)+1))
		args = append(args, "%"+text+"%")
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += " LIMIT $" + strconv.Itoa(len(args)+1)
	args = append(args, limit)

	query += " OFFSET $" + strconv.Itoa(len(args)+1)
	args = append(args, offset)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := models.Song{}
		err = rows.Scan(&s.ID, &s.GroupName, &s.SongName, &s.ReleaseDate, &s.Text, &s.Link)
		if err != nil {
			return nil, err
		}
		songs = append(songs, s)
	}

	return songs, nil
}

func GetDetailSong(id string) (*models.Song, error) {
	rows, err := db.DB.Query(`SELECT id, group_name, song_name, release_date, text, link FROM songs WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	s := models.Song{}
	err = rows.Scan(&s.ID, &s.GroupName, &s.SongName, &s.ReleaseDate, &s.ReleaseDate, &s.Text, &s.Link)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func PostSong(s models.Song) error {
	_, err := db.DB.Exec(`INSERT INTO songs(group_name, song_name, release_date, text, link) VALUES ($1, $2, $3, $4, $5)`,
		s.GroupName, s.SongName, s.ReleaseDate, s.ReleaseDate, s.Text, s.Link)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSong(id string, s models.Song) error {
	_, err := db.DB.Exec(`UPDATE songs SET group_name=$1, song_name=$2, release_date=$3, text=$4, link=$5, WHERE id=$6`,
		s.GroupName, s.SongName, s.ReleaseDate, s.ReleaseDate, s.Text, s.Link, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSong(id string) error {
	_, err := db.DB.Query(`DELETE FROM songs WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
