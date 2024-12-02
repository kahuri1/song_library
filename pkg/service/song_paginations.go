package service

import (
	"errors"
	"github.com/kahuri1/song_library/pkg/model"
	"strings"
)

func (s *Service) SongLine(song *model.SongPaginations) (*model.SongPaginations, error) {
	text, err := s.repo.SongText(song.SongID)
	if err != nil {
		return nil, errors.New("failed to fetch from DB: " + err.Error())
	}
	song.Text = text

	err = s.PaginationTextSong(song)
	if err != nil {
		return nil, errors.New("failed pagination text: " + err.Error())
	}

	return song, nil
}

func (s *Service) PaginationTextSong(song *model.SongPaginations) error {
	song.Lines = strings.Split(song.Text, "\n")

	start := (song.Page - 1) * song.Limit
	end := start + song.Limit
	var filteredLines []string
	for _, line := range song.Lines {
		if line != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	song.Lines = filteredLines
	if start >= len(song.Lines) {
		return errors.New("zero array")
	}
	if end > len(song.Lines) {
		end = len(song.Lines)
	}
	song.Lines = song.Lines[start:end]
	return nil
}
