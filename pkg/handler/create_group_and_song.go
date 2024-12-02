package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Create group and song
// @Tags Library
// @Description Allows you to create a group and a song in one request
// @ID create-group-and-song
// @Accept json
// @Produce json
// @Param input body model.Input true "Input data for creating group and song"
// @Success 201 {object} model.Response "Group and song created successfully"
// @Failure 400 {object} model.Error "Invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /group/song [post]
func (h *Handler) CreateGroupAndSong(c *gin.Context) {
	var input model.Input

	d, err := c.GetRawData()
	err = json.Unmarshal(d, &input)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}
	err = h.service.CreateGroupAndSong(&input)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}
	sendResponse(c, http.StatusCreated, "group and song created", nil)
}
