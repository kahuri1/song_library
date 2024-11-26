package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) UpdateSong(query string, args []interface{}, input *model.Input) (*model.Input, error) {

	if len(args) != 0 {
		_, err := r.db.Exec(query, args...)
		if err != nil {
			return nil, fmt.Errorf("failed to update song: %w", err)
		}
	}

	input, err := r.GetSong(input)
	if err != nil {
		return nil, fmt.Errorf("failed get song: %w", err)
	}
	return input, nil
}

func (r *Repository) GetSong(input *model.Input) (*model.Input, error) {
	checkSql, checkArgs, err := sq.
		Select("title, group_id, release_date, lyrics, link ").
		From("songs").
		Where("song_id = ?", input.Song.SongID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err // Возвращаем ошибку, если не удалось создать SQL-запрос
	}
	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&input.Song.Title, &input.Song.GroupID, &input.Song.ReleaseDate, &input.Song.Lyrics, &input.Song.Link)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Возвращаем nil, если группы не найдены
		}
		return nil, err // Возвращаем ошибку, если что-то пошло не так при выполнении запроса
	}
	return input, nil
}
