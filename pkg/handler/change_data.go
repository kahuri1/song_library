package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

func (h *Handler) ChangeData(c *gin.Context) {
	var input model.Input

	d, err := c.GetRawData()
	err = json.Unmarshal(d, &input)
	if err != nil {
		handlerError(c, err, "ошибка считывания json", http.StatusBadRequest)
		return
	}

	ResponseServise, err := h.service.ChangeData(&input)
	if err != nil {
		handlerError(c, err, "ошибка обработки запроса", http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"данные изменены": ResponseServise})

}
