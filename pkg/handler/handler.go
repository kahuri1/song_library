package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	log "github.com/sirupsen/logrus"
)

type songLibsService interface {
	CreateGroup(group *model.Group) error
	CreateSongAndDetails(song *model.Song) error
	CreateGroupAndSong(input *model.Input) error
}

type Handler struct {
	service songLibsService
}

func Newhandler(service songLibsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/group", h.CreateGroup)
	r.POST("/group/song", h.CreateGroupAndSong)
	r.POST("/song", h.CreateSongAndDetails)
	return r
}

func handlerError(c *gin.Context, err error, message string, statusCode int) {
	log.WithFields(log.Fields{
		"error":   err.Error(),
		"context": c.Request.URL.Path,
	}).Error(message) // Логирование ошибки

	c.JSON(statusCode, gin.H{"error": message})
}
