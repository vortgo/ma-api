package handler

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"ma-api/presenter"
	"net/http"
	"strconv"
)

func BandById(context echo.Context) error {
	band := models.Band{}
	bandId, _ := strconv.Atoi(context.Param("id"))
	db := repositories.PostgresDB.
		Preload("Genres").
		Preload("Label").
		Preload("Country").
		Preload("LyricalThemes").Find(&band, bandId)

	if db.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	simpleBand := presenter.NewBandPresenter(&band).SimpleBand()

	return context.JSON(200, &simpleBand)
}

func BandAlbums(context echo.Context) error {
	var albums []*models.Album
	bandId, _ := strconv.Atoi(context.Param("id"))
	repositories.PostgresDB.Preload("Label").Where(&models.Album{BandID: uint(bandId)}).Find(&albums)

	albumsCollection := presenter.NewBandAlbumsPresenter(albums).AlbumsCollection()

	return context.JSON(200, &albumsCollection)
}
