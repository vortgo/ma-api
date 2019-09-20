package handler

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"ma-api/presenter"
	"ma-api/services"
	"net/http"
	"os"
	"strconv"
)

func SongById(c echo.Context) error {
	song := models.Song{}
	songId, _ := strconv.Atoi(c.Param("id"))
	db := repositories.PostgresDB.Preload("Album").Preload("Band").First(&song, songId)

	if db.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	q := song.Band.Name + " - " + song.Name
	youtubeSearcher := services.NewYoutubeSearcher(os.Getenv("YOUTUBE_KEY"))
	id, title := youtubeSearcher.Search(q)

	var videoKey string

	if youtubeSearcher.ValidateResult(song.Name, title) {
		videoKey = id
	}

	simpleSong := presenter.NewSongPresenter(&song).SimpleSongs(videoKey)

	return c.JSON(200, &simpleSong)
}
