package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Create group
// @Tags Group
// @Description create group
// @ID create-group
// @Accept json
// @Produce json
// @Param input body model.Input true "Input data for updating song"
// @Success 200 {object} model.Response "The group was successfully created"
// @Failure 400 {object} model.Error "Invalid input data"
// @Failure 404 {object} model.Error "Song not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /group [post]
func (h *Handler) CreateGroup(c *gin.Context) {
	var group model.Group

	d, err := c.GetRawData()

	err = json.Unmarshal(d, &group)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	err = h.service.CreateGroup(&group)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}
	sendResponse(c, http.StatusOK, "group created", nil)
}
