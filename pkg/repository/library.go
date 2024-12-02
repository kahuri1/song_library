package repository

import (
	"github.com/kahuri1/song_library/pkg/model"
)

func (r *Repository) Library(query string) (*model.Library, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var library model.Library
	for rows.Next() {
		var input model.Input // Предполагается, что Song - это структура, содержащая данные о песне
		// Предполагается, что Group - структура для группы

		// Заполнение данных
		if err := rows.Scan(&input.Song.SongID, &input.Song.Title, &input.Song.ReleaseDate, &input.Song.Lyrics, &input.Song.Link, &input.Group.ID, &input.Group.Name); err != nil {
			return nil, err
		}
		library.Library = append(library.Library, input)
		// Добавление группы, если нужно.
	}

	return &library, nil
}
