package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Filter Library
// @Tags Library
// @Description Allows you to filter the library based on the provided criteria
// @ID filter-library
// @Accept json
// @Produce json
// @Param input body model.LibraryRequest true "Input data for filtering the library"
// @Success 200 {object} model.Response "Filtered library data"
// @Failure 400 {object} model.Error "Invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /library [post]
func (h *Handler) Library(c *gin.Context) {
	var LibraryRequest model.LibraryRequest
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &LibraryRequest)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	Library, err := h.service.Library(&LibraryRequest)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}

	sendResponse(c, http.StatusOK, "libraryFilter", Library)
}
