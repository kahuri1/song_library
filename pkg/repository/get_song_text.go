package repository

import (
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) SongText(songID int64) (string, error) {
	var text string
	checkSql, checkArgs, err := sq.
		Select("lyrics").
		From("songs").
		Where("song_id = ?", songID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", fmt.Errorf("failed to check query songs: %w", err)
	}
	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&text)
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("failed to check query songs: %w", err)
	}

	if len(text) != 0 {
		return text, nil
	}
	return "", errors.New("text songs zero: " + err.Error())
}
