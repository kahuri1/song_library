package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) UpdateGroup(query string, args []interface{}, input *model.Input) (*model.Input, error) {

	if input.Group.Name != "" {
		_, err := r.db.Exec(query, args...)
		if err != nil {
			return nil, fmt.Errorf("failed to update group: %w", err)
		}
	}

	input, err := r.GetGroup(input)
	if err != nil {
		return nil, fmt.Errorf("failed get group: %w", err)
	}
	return input, nil
}

func (r *Repository) GetGroup(input *model.Input) (*model.Input, error) {
	checkSql, checkArgs, err := sq.
		Select("name").
		From("groups").
		Where("group_id = ?", input.Group.ID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}
	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&input.Group.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return input, nil
}
