package presenter

import (
	"github.com/vortgo/ma-parser/models"
	"time"
)

type collectionLatestBandsUpdatePresenter struct {
	latestBands []*models.LatestBandUpdate
}

type latestBandsUpdateSimpleCollection struct {
	ID   int    `json:"id"`
	Name string `json:"title"`
	Date string `json:"date"`
}

func NewCollectionLatestBandsUpdatePresenter(latestBands []*models.LatestBandUpdate) *collectionLatestBandsUpdatePresenter {
	return &collectionLatestBandsUpdatePresenter{latestBands: latestBands}
}

func (presenter *collectionLatestBandsUpdatePresenter) SimpleCollection() []*latestBandsUpdateSimpleCollection {
	var collection []*latestBandsUpdateSimpleCollection
	for _, latestBand := range presenter.latestBands {
		collection = append(collection, &latestBandsUpdateSimpleCollection{
			ID:   int(latestBand.BandID),
			Name: latestBand.Band.Name,
			Date: latestBand.UpdatedAt.Format(time.RFC3339),
		})
	}

	return collection
}
