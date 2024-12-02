package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Create song and details
// @Tags Song
// @Description Allows you to create a song and its related details in one request
// @ID create-song-and-details
// @Accept json
// @Produce json
// @Param input body model.Song true "Input data for creating song and its details"
// @Success 201 {object} model.Response "Song created successfully"
// @Failure 400 {object} model.Error "Invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /song [post]
func (h *Handler) CreateSongAndDetails(c *gin.Context) {
	var song model.Song

	d, err := c.GetRawData()
	err = json.Unmarshal(d, &song)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	err = h.service.CreateSongAndDetails(&song)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}

	sendResponse(c, http.StatusCreated, "Song created successfully", nil)
}
