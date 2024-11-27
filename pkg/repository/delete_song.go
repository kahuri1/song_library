package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) DeleteSongByID(song *model.Song) error {
	sql, args, err := sq.
		Delete("songs").
		Where("song_id = ?", song.SongID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteSongByName(song *model.Song) error {
	sql, args, err := sq.
		Delete("songs").
		Where("title = ?", song.Title).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}
