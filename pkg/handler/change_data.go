package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Change Data
// @Tags Library
// @Description Update the details of a song and group
// @ID update-song-and-group
// @Accept json
// @Produce json
// @Param input body model.Input true "Input data for updating song"
// @Success 200 {object} model.Response "Song data updated successfully"
// @Failure 400 {object} model.Error "Invalid input data"
// @Failure 404 {object} model.Error "Song not found"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /group/song [put]
func (h *Handler) ChangeData(c *gin.Context) {
	var input model.Input

	d, err := c.GetRawData()
	err = json.Unmarshal(d, &input)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	ResponseServise, err := h.service.ChangeData(&input)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}

	sendResponse(c, http.StatusOK, "data", ResponseServise)

}
