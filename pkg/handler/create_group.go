package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

func (h *Handler) CreateGroup(c *gin.Context) {
	var group model.Group

	d, err := c.GetRawData()

	err = json.Unmarshal(d, &group)
	if err != nil {
		handlerError(c, err, "ошибка считывания json", http.StatusBadRequest)
		return
	}

	err = h.service.CreateGroup(&group)
	if err != nil {
		handlerError(c, err, "ошибка обработки запроса", http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "group created"})
}
