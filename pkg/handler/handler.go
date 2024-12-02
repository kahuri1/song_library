package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kahuri1/song_library/docs"
	"github.com/kahuri1/song_library/pkg/model"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type songLibsService interface {
	CreateGroup(group *model.Group) error
	CreateSongAndDetails(song *model.Song) error
	CreateGroupAndSong(input *model.Input) error
	ChangeData(input *model.Input) (*model.Input, error)
	DeleteGroup(group *model.Group) error
	DeleteSong(song *model.Song) error
	Library(LibraryRequest *model.LibraryRequest) (*model.Library, error)
	SongLine(song *model.SongPaginations) (*model.SongPaginations, error)
}

type Handler struct {
	service songLibsService
}

func Newhandler(service songLibsService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/group", h.CreateGroup)             //
	r.POST("/group/song", h.CreateGroupAndSong) //
	r.PUT("/group/song", h.ChangeData)          //
	r.POST("/song", h.CreateSongAndDetails)     //
	r.POST("/song/text", h.SongLine)            //
	r.DELETE("/group", h.DeleteGroup)
	r.DELETE("/song", h.DeleteSong)
	r.POST("/songs", h.NewSong)   //
	r.POST("/library", h.Library) //
	return r
}

func handlerError(c *gin.Context, err error, message string, statusCode int) {
	errorResponse := &model.Error{
		Code:    statusCode,
		Message: message,
	}

	log.WithFields(log.Fields{
		"error":   err.Error(),
		"context": c.Request.URL.Path,
	}).Error(message)

	c.JSON(statusCode, errorResponse)
}

func sendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	response := model.Response{
		Status:  http.StatusText(statusCode),
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}
