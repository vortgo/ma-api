package presenter

import "github.com/vortgo/ma-parser/models"

type omit *struct{}

type albumInCollection struct {
	*simpleAlbum
	BandName omit `json:"band_name,omitempty"`
}

type bandAlbumsPresenter struct {
	albums []*models.Album
}

func NewBandAlbumsPresenter(albums []*models.Album) *bandAlbumsPresenter {
	return &bandAlbumsPresenter{albums: albums}
}

func (presenter *bandAlbumsPresenter) AlbumsCollection() []*albumInCollection {
	var collection []*albumInCollection

	for _, album := range presenter.albums {
		albumInCollection := albumInCollection{simpleAlbum: NewAlbumPresenter(album).SimpleAlbum()}
		collection = append(collection, &albumInCollection)
	}

	return collection
}
