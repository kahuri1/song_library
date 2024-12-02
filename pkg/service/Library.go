package service

import (
	"errors"
	"fmt"
	"github.com/kahuri1/song_library/pkg/model"
	"strings"
)

func (s *Service) Library(LibraryRequest *model.LibraryRequest) (*model.Library, error) {
	query := `
		SELECT  s.song_id, s.title, s.release_date, s.lyrics, s.link, g.group_id , g.name
		FROM songs s 
		LEFT JOIN "groups" g ON s.group_id = g.group_id 
		WHERE 1=1`

	var filter []string
	if LibraryRequest.Filters.Title != "" {
		query += fmt.Sprintf(` AND s.title LIKE  '%s'`, "%"+LibraryRequest.Filters.Title+"%")
	}

	if LibraryRequest.Filters.ReleaseDate != "" {
		query += fmt.Sprintf(` AND s.release_date = '%s'`, LibraryRequest.Filters.ReleaseDate)
	}
	if LibraryRequest.Filters.GroupName != "" {
		query += fmt.Sprintf(` AND g.name LIKE '%s'`, "%"+LibraryRequest.Filters.GroupName+"%")
	}
	if LibraryRequest.Filters.Text != "" {
		query += fmt.Sprintf(` AND s.lyrics LIKE '%s'`, "%"+LibraryRequest.Filters.Text+"%")
	}

	pageSize := LibraryRequest.Pagination.PageSize
	page := LibraryRequest.Pagination.Page
	offset := (page - 1) * pageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	query += strings.Join(filter, " ")

	library, err := s.repo.Library(query)
	if err != nil {
		return nil, errors.New("failed to fetch from DB: " + err.Error())
	}
	return library, nil
}
