package handler

import (
	"github.com/labstack/echo"
	"ma-api/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/stat", ApiStat)

	e.GET("/song-iframe", SongIframe)
	api := e.Group("/api/v1")
	api.Use(middleware.RequestStatisticMiddleware)

	band := api.Group("/band")
	band.GET("/:id", BandById)
	band.GET("/:id/albums", BandAlbums)
	band.GET("/:id/reviews", BandReviews)

	album := api.Group("/album")
	album.GET("/:id", AlbumById)
	album.GET("/:id/songs", AlbumSongs)
	album.GET("/:id/reviews", AlbumReviews)

	song := api.Group("/song")
	song.GET("/:id", SongById)

	collection := api.Group("/collection")
	collection.GET("/latest-bands-update", LatestBandsUpdate)
	collection.GET("/upcoming-albums", UpcomingAlbums)

	review := api.Group("/review")
	review.GET("/:id", Review)

	api.GET("/search", Search)

}
