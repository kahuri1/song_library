package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

func (h *Handler) DeleteSong(c *gin.Context) {
	var song model.Song
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &song)
	if err != nil {
		handlerError(c, err, "ошибка считывания json", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteSong(&song)
	if err != nil {
		handlerError(c, err, "ошибка обработки запроса", http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "song delete"})
}
