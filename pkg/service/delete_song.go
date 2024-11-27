package service

import (
	"errors"
	"github.com/kahuri1/song_library/pkg/model"
)

func (s *Service) DeleteSong(song *model.Song) error {
	if song.SongID != 0 {
		_, err := s.repo.CheckSong(song)
		if err == nil {
			return errors.New("the group does not exist")
		}

		err = s.repo.DeleteSongByID(song)
		if err != nil {
			return errors.New("error deleting song")
		}
		return nil
	} else if song.Title != "" {
		err := s.repo.DeleteSongByName(song)
		if err != nil {
			return errors.New("error deleting song")
		}
		return nil
	}
	return errors.New("zero json")
}
