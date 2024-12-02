package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Delete song
// @Tags Song
// @Description Allows you to delete a song by its ID
// @ID delete-song
// @Accept json
// @Produce json
// @Param input body model.Song true "Input data for deleting a song"
// @Success 200 {object} model.Response "song deleted"
// @Failure 400 {object} model.Error "Invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /song [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	var song model.Song
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &song)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteSong(&song)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}

	sendResponse(c, http.StatusOK, "song delete", nil)
}
