package presenter

import (
	"github.com/vortgo/ma-parser/models"
	"ma-api/utils"
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

		image := upcomingAlbum.Album.Image
		if image == "" {
			image = utils.MakeAppUrl("static/empty-album.jpg")
		}

		collection = append(collection, &upcomingAlbumsSimpleCollection{
			ID:       int(upcomingAlbum.AlbumID),
			Name:     upcomingAlbum.Album.Name,
			BandName: upcomingAlbum.Album.Band.Name,
			Image:    image,
		})
	}

	return collection
}
