package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type songLibsService interface {
}

type Handler struct {
	service songLibsService
}

func Newhandler(service songLibsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	return r
}

func handlerError(c *gin.Context, err error, message string, statusCode int) {
	log.WithFields(log.Fields{
		"error":   err.Error(),
		"context": c.Request.URL.Path,
	}).Error(message) // Логирование ошибки

	c.JSON(statusCode, gin.H{"error": message})
}
