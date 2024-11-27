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

func (h *Handler) NewSong(c *gin.Context) {
	var song model.AddedSong
	var input model.Input
	var songDetail model.SongsDetail

	d, err := c.GetRawData()
	err = json.Unmarshal(d, &song)
	if err != nil {
		handlerError(c, err, "ошибка считывания json", http.StatusBadRequest)
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
		handlerError(c, err, "ошибка обработки ответа", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &songDetail); err != nil {
		handlerError(c, err, "ошибка считывания json", http.StatusBadRequest)
		return
	}
	input.Song.Link = songDetail.Link
	input.Song.ReleaseDate = songDetail.ReleaseDate
	input.Song.Lyrics = songDetail.Text

	err = h.service.CreateGroupAndSong(&input)
	if err != nil {
		handlerError(c, err, "ошибка обработки запроса", http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "group and song created"})

}
