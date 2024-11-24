package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) CreateSongAndDetails(song *model.Song) error {

	sql, args, err := sq.
		Insert("songs").
		Columns("title", "group_id", "release_date", "lyrics", "link").
		Values(song.Title, song.GroupID, song.ReleaseDate, song.Lyrics, song.Link).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to create song creation request: %w", err)
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("failed to process song creation query: %w", err)
	}
	return nil
}
