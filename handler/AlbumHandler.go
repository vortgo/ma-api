package handler

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"ma-api/presenter"
)

func AlbumById(context echo.Context) error {
	album := models.Album{}
	repositories.PostgresDB.Preload("Label").Preload("Band").First(&album)

	albumPresenter := presenter.NewAlbumPresenter(&album)
	albumPresenter.SimpleAlbum().ToJson()

	context.String(200, albumPresenter.SimpleAlbum().WithBandInfo().ToJson())

	return nil
}

func AlbumSongs(context echo.Context) error {

	return nil
}
