package service

import (
	"github.com/kahuri1/song_library/pkg/model"
	log "github.com/sirupsen/logrus"
)

type repo interface {
	CheckGroup(group *model.Group) (int64, error)
	CreateGroup(group *model.Group) error
	CheckSong(song *model.Song) (int64, error)
	CreateSongAndDetails(song *model.Song) error
	UpdateGroup(queryGroup string, paramGroup []interface{}, input *model.Input) (*model.Input, error)
	UpdateSong(query string, args []interface{}, input *model.Input) (*model.Input, error)
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	log.Info("service init")

	return &Service{
		repo: repo,
	}
}
