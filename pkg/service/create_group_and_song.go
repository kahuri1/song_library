package service

import (
	"github.com/kahuri1/song_library/pkg/model"
)

func (s *Service) CreateGroupAndSong(input *model.Input) error {
	err := s.CreateGroup(&input.Group)
	if err != nil {
		return err
	}
	idGroup, err := s.repo.CheckGroup(&input.Group)
	if idGroup != 0 {
		input.Song.GroupID = idGroup
	}

	err = s.CreateSongAndDetails(&input.Song)
	if err != nil {
		return err
	}
	return nil
}
