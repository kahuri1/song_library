package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"net/http"
)

// @Summary Delete group
// @Tags Group
// @Description Allows you to delete a group by its ID
// @ID delete-group
// @Accept json
// @Produce json
// @Param input body model.Group true "Input data for deleting a group"
// @Success 200 {object} model.Response "group deleted"
// @Failure 400 {object} model.Error "Invalid input data"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /group [delete]
func (h *Handler) DeleteGroup(c *gin.Context) {
	var group model.Group
	d, err := c.GetRawData()
	err = json.Unmarshal(d, &group)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteGroup(&group)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}

	sendResponse(c, http.StatusOK, "group deleted", nil)
}
