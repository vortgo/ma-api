package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"ma-api/presenter"
)

func main() {
	//e := echo.New()
	//handler.RegisterRoutes(e)
	//e.Logger.Fatal(e.Start(":1323"))

	album := models.Album{}
	repositories.PostgresDB.Preload("Label").Preload("Band").First(&album)

	albumPresenter := presenter.NewAlbumPresenter(&album)
	fmt.Println(albumPresenter.SimpleAlbum().WithBandInfo().ToJson())
}
