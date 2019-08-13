package presenter

import (
	"encoding/json"
	"fmt"
	"github.com/vortgo/ma-parser/models"
	"time"
)

type source interface {
}

type simpleAlbumPresent struct {
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

type AlbumPresenter struct {
	album  *models.Album
	source interface{}
}

func NewAlbumPresenter(album *models.Album) *AlbumPresenter {
	return &AlbumPresenter{album: album, source: album}
}

func (presenter *AlbumPresenter) SimpleAlbum() *AlbumPresenter {
	simpleAlbum := struct {
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
	}{
		Id:          int(presenter.album.ID),
		BandId:      int(presenter.album.BandID),
		Type:        presenter.album.Type,
		Name:        presenter.album.Name,
		BandName:    presenter.album.Band.Name,
		Year:        presenter.album.Year,
		ReleaseDate: presenter.album.ReleaseDate,
		Label:       presenter.album.Label.Name,
		Image:       presenter.album.Image,
		TotalTime:   presenter.album.TotalTime,
	}
	presenter.source = simpleAlbum
	return presenter
}

func (presenter *AlbumPresenter) WithBandInfo() *AlbumPresenter {
	albumWithBandInfo := struct {
		*source
		BandCountry string `json:"band_country"`
	}{
		source:      presenter.source.(source),
		BandCountry: "123",
	}

	presenter.source = albumWithBandInfo
	return presenter
}

func (presenter *AlbumPresenter) ToJson() string {
	fmt.Println(presenter.source)
	b, _ := json.Marshal(presenter.source)
	return string(b)
}
