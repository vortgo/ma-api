package handler

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/olivere/elastic/v7"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"github.com/vortgo/ma-parser/utils"
	"log"
	"ma-api/presenter"
	"os"
)

type searchEntityData struct {
	Name string `json:"name"`
	Data interface{}
}

func Search(context echo.Context) error {
	findPhrase := context.QueryParam("search")
	var result []*searchEntityData
	var (
		bands  []*models.Band
		albums []*models.Album
	)

	var (
		band  = models.Band{}
		album = models.Album{}
	)

	bandIds := searchIds(band.GetIndexName(), findPhrase)

	repositories.PostgresDB.
		Where(`id in (?)`, bandIds).
		Order("id asc").
		Find(&bands)

	albumsIds := searchIds(album.GetIndexName(), findPhrase)

	repositories.PostgresDB.
		Preload("Band").
		Where(`id in (?)`, albumsIds).
		Order("id asc").
		Find(&albums)

	bandsFormatted := presenter.NewCollectionBandPresenter(bands).SearchBands()
	albumsFormatted := presenter.NewCollectionAlbumPresenter(albums).SearchAlbum()

	result = append(result, &searchEntityData{Name: "Bands", Data: bandsFormatted})
	result = append(result, &searchEntityData{Name: "Albums", Data: albumsFormatted})

	return context.JSON(200, &result)
}

func searchIds(index, search string) []interface{} {
	var ids []interface{}
	es, err := elastic.NewClient(elastic.SetHttpClient(utils.CustomHttpClient), elastic.SetSniff(false), elastic.SetHealthcheck(false), elastic.SetURL(os.Getenv("ELASTIC_URL")))
	if err != nil {
		log.Printf("Elastic: %s\n", err)
		return ids
	}

	ctx := context.Background()
	query := elastic.NewMatchQuery("name", search)
	searchResult, err := es.Search().
		Index(index).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)

	for _, hit := range searchResult.Hits.Hits {
		item := make(map[string]interface{})
		json.Unmarshal(hit.Source, &item)
		ids = append(ids, item[`id`])
	}

	return ids
}
