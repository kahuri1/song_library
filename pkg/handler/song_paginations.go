package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Retrieve Song Lines
// @Tags Song
// @Description This endpoint retrieves the lines/lyrics of a song based on the provided song ID.
// @ID get-song-lines
// @Accept json
// @Produce json
// @Param input body model.SongPaginations true "Input data containing the song ID"
// @Success 200 {object} model.Response "The lyrics of the requested song"
// @Failure 400 {object} model.Error "Invalid input data or errors retrieving song lines"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /song/text [post]
func (h *Handler) SongLine(c *gin.Context) {
	var songID model.SongPaginations
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &songID)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	song, err := h.service.SongLine(&songID)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}
	sendResponse(c, http.StatusOK, "text", song.Lines)
}
