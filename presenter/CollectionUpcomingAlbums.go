package presenter

import (
	"github.com/vortgo/ma-parser/models"
)

type collectionUpcomingAlbumsPresenter struct {
	upcomingAlbums []*models.UpcomingAlbum
}

type upcomingAlbumsSimpleCollection struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	BandName string `json:"band_name"`
	Image    string `json:"image"`
}

func NewCollectionUpcomingAlbumsPresenter(upcomingAlbums []*models.UpcomingAlbum) *collectionUpcomingAlbumsPresenter {
	return &collectionUpcomingAlbumsPresenter{upcomingAlbums: upcomingAlbums}
}

func (presenter *collectionUpcomingAlbumsPresenter) SimpleCollection() []*upcomingAlbumsSimpleCollection {
	var collection []*upcomingAlbumsSimpleCollection
	for _, upcomingAlbum := range presenter.upcomingAlbums {
		collection = append(collection, &upcomingAlbumsSimpleCollection{
			ID:       int(upcomingAlbum.AlbumID),
			Name:     upcomingAlbum.Album.Name,
			BandName: upcomingAlbum.Album.Band.Name,
			Image:    upcomingAlbum.Album.Image,
		})
	}

	return collection
}
