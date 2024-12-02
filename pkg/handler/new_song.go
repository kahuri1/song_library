package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/song_library/pkg/model"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

// @Summary Add New Song
// @Tags Library
// @Description This endpoint allows you to add a new song by retrieving its details from an external API based on the provided group and song name.
// @ID add-new-song
// @Accept json
// @Produce json
// @Param input body model.AddedSong true "Input data for adding a new song"
// @Success 200 {object} model.Response "Success message indicating the song was created"
// @Failure 400 {object} model.Error "Invalid input data or errors communicating with external API"
// @Failure 500 {object} model.Error "Internal server error"
// @Router /songs [post]
func (h *Handler) NewSong(c *gin.Context) {
	var song model.AddedSong
	var input model.Input
	var songDetail model.SongsDetail

	d, err := c.GetRawData()
	err = json.Unmarshal(d, &song)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}
	externalUrl := viper.GetString("externalUrl")

	groupString := url.QueryEscape(song.Group)
	songString := url.QueryEscape(song.Song)

	api := fmt.Sprintf("%s?group=%s&song=%s", externalUrl, groupString, songString)

	resp, err := http.Get(api)
	if err != nil {
		handlerError(c, err, "error getting data from external api", http.StatusBadRequest)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		handlerError(c, err, "status code is different from 200", http.StatusBadRequest)
		return
	}
	input.Group.Name = song.Group
	input.Song.Title = song.Song

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		handlerError(c, err, "error read json", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &songDetail); err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}
	input.Song.Link = songDetail.Link
	input.Song.ReleaseDate = songDetail.ReleaseDate
	input.Song.Lyrics = songDetail.Text

	err = h.service.CreateGroupAndSong(&input)
	if err != nil {
		handlerError(c, err, "Error processing request", http.StatusBadRequest)
		return
	}

	sendResponse(c, http.StatusCreated, "group and song created", nil)
}
