package presenter

import (
	"github.com/vortgo/ma-parser/models"
	"ma-api/utils"
	"time"
)

type simpleAlbum struct {
	Id          int       `json:"id"`
	BandId      int       `json:"band_id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	BandName    string    `json:"band_name"`
	Year        int       `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Label       string    `json:"label"`
	Image       string    `json:"image"`
	TotalTime   string    `json:"total_time"`
}

type searchAlbum struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	BandName string `json:"band_name"`
}

type albumPresenter struct {
	album *models.Album
}

type collectionAlbumPresenter struct {
	albums []*models.Album
}

func NewAlbumPresenter(album *models.Album) *albumPresenter {
	return &albumPresenter{album: album}
}

func NewCollectionAlbumPresenter(albums []*models.Album) *collectionAlbumPresenter {
	return &collectionAlbumPresenter{albums: albums}
}

func (presenter *albumPresenter) SimpleAlbum() *simpleAlbum {
	image := presenter.album.Image
	if image == "" {
		image = utils.MakeAppUrl("static/empty-album.jpg")
	}

	return &simpleAlbum{
		Id:          int(presenter.album.ID),
		BandId:      int(presenter.album.BandID),
		Type:        presenter.album.Type,
		Name:        presenter.album.Name,
		BandName:    presenter.album.Band.Name,
		Year:        presenter.album.Year,
		ReleaseDate: presenter.album.ReleaseDate,
		Label:       presenter.album.Label.Name,
		Image:       image,
		TotalTime:   presenter.album.TotalTime,
	}
}

func (presenter *collectionAlbumPresenter) SearchAlbum() []*searchAlbum {
	var collection []*searchAlbum

	for _, album := range presenter.albums {
		collection = append(collection, &searchAlbum{
			Id:       int(album.ID),
			Name:     album.Name,
			BandName: album.Band.Name,
		})
	}

	return collection
}
