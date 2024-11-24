package service

import (
	"errors"
	"github.com/kahuri1/song_library/pkg/model"
)

func (s *Service) CreateGroup(group *model.Group) error {
	idGroup, err := s.repo.CheckGroup(group)
	if idGroup != 0 {
		return errors.New("the group was created earlier")
	}
	err = s.repo.CreateGroup(group)
	if err != nil {
		return err
	}
	return nil
}
