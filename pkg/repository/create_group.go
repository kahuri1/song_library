package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) CreateGroup(group *model.Group) error {

	sql, args, err := sq.
		Insert("groups").
		Columns("name").
		Values(group.Name).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to create name group creation request: %w", err)
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("failed to process group creation query: %w", err)
	}
	return nil
}
