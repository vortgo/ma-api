package handler

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"ma-api/presenter"
)

func LatestBandsUpdate(context echo.Context) error {
	var latestBandsUpdate []*models.LatestBandUpdate
	repositories.PostgresDB.Preload("Band").Limit(10).Order("updated_at desc").Find(&latestBandsUpdate)

	collection := presenter.NewCollectionLatestBandsUpdatePresenter(latestBandsUpdate).SimpleCollection()

	return context.JSON(200, &collection)
}

func UpcomingAlbums(context echo.Context) error {
	var upcomingAlbums []*models.UpcomingAlbum
	repositories.PostgresDB.Preload("Album.Band").Limit(10).Order("updated_at desc").Find(&upcomingAlbums)

	collection := presenter.NewCollectionUpcomingAlbumsPresenter(upcomingAlbums).SimpleCollection()

	return context.JSON(200, collection)
}
