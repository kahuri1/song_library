package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) DeleteGroupByID(group *model.Group) error {
	sql, args, err := sq.
		Delete("groups").
		Where("group_id = ?", group.ID).
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

func (r *Repository) DeleteGroupByName(group *model.Group) error {
	sql, args, err := sq.
		Delete("groups").
		Where("name = ?", group.Name).
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
