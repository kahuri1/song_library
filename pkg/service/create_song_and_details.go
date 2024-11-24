package service

import (
	"errors"
	"github.com/kahuri1/song_library/pkg/model"
)

func (s *Service) CreateSongAndDetails(song *model.Song) error {
	idGroup, err := s.repo.CheckSong(song)
	if idGroup != 0 {
		return errors.New("the song was created earlier")
	}

	err = s.repo.CreateSongAndDetails(song)
	if err != nil {
		return err
	}

	return nil
}
