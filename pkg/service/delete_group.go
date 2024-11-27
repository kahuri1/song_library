package service

import (
	"errors"
	"github.com/kahuri1/song_library/pkg/model"
)

func (s *Service) DeleteGroup(group *model.Group) error {
	if group.ID != 0 {
		_, err := s.repo.CheckGroup(group)
		if err == nil {
			return errors.New("the group does not exist")
		}

		err = s.repo.DeleteGroupByID(group)
		if err != nil {
			return errors.New("error deleting group")
		}

		return nil
	} else if group.Name != "" {
		err := s.repo.DeleteGroupByName(group)
		if err != nil {
			return errors.New("error deleting group")
		}

		return nil
	}
	return errors.New("zero json")
}
