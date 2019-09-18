package handler

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"ma-api/presenter"
	"net/http"
	"strconv"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func AlbumById(c echo.Context) error {
	album := models.Album{}
	albumId, _ := strconv.Atoi(c.Param("id"))
	db := repositories.PostgresDB.Preload("Label").Preload("Band").First(&album, albumId)

	if db.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	simpleAlbum := presenter.NewAlbumPresenter(&album).SimpleAlbum()

	return c.JSON(200, &simpleAlbum)
}

func AlbumSongs(c echo.Context) error {
	var songs []*models.Song
	albumId, _ := strconv.Atoi(c.Param("id"))
	repositories.PostgresDB.Where(&models.Song{AlbumID: uint(albumId)}).Find(&songs)

	songsCollection := presenter.NewAlbumSongsPresenter(songs).SongsCollection()
	return c.JSON(200, &songsCollection)
}
