package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) CheckGroup(group *model.Group) (int64, error) {
	var id int64

	checkSql, checkArgs, err := sq.
		Select("group_id").
		From("groups").
		Where("name = ?", group.Name).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create check query: %w", err)
	}

	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to execute check for existing group name: %w", err)
	}

	// Если группа существует, возвращаем ее id
	if id != 0 {
		return id, fmt.Errorf("the group has already been created")
	}
	return 0, nil
}
