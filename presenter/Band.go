package presenter

import (
	"github.com/vortgo/ma-parser/models"
	"ma-api/utils"
	"strconv"
)

type bandPresenter struct {
	band *models.Band
}

type collectionBandPresenter struct {
	bands []*models.Band
}

type searchBand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type simpleBand struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Country      string   `json:"country"`
	Status       string   `json:"status"`
	FormedIn     int      `json:"formed_in"`
	Label        string   `json:"label"`
	Description  string   `json:"description"`
	ImageLogo    string   `json:"image_logo"`
	ImageBand    string   `json:"image_band"`
	Genres       []string `json:"genres"`
	LyricalTheme []string `json:"lyrical_theme"`
}

func NewCollectionBandPresenter(bands []*models.Band) *collectionBandPresenter {
	return &collectionBandPresenter{bands: bands}
}

func NewBandPresenter(band *models.Band) *bandPresenter {
	return &bandPresenter{band: band}
}

func (presenter *bandPresenter) SimpleBand() *simpleBand {
	var genres []string
	var lyricalTheme []string

	for _, genre := range presenter.band.Genres {
		genres = append(genres, genre.Name)
	}

	for _, bandLyricalTheme := range presenter.band.LyricalThemes {
		lyricalTheme = append(lyricalTheme, bandLyricalTheme.Name)
	}

	imageBand := presenter.band.ImageBand
	if imageBand == "" {
		imageBand = utils.GetBandStabImageUrl(int(presenter.band.ID))
	}

	return &simpleBand{
		ID:           int(presenter.band.ID),
		Name:         presenter.band.Name,
		Country:      presenter.band.Country.Name,
		Status:       presenter.band.Status,
		FormedIn:     presenter.band.FormedIn,
		Label:        presenter.band.Label.Name,
		Description:  presenter.band.Description,
		ImageLogo:    presenter.band.ImageLogo,
		ImageBand:    imageBand,
		Genres:       genres,
		LyricalTheme: lyricalTheme,
	}
}

func (presenter *collectionBandPresenter) SearchBands() []*searchBand {
	var collection []*searchBand

	for _, band := range presenter.bands {
		collection = append(collection, &searchBand{
			ID:   int(band.ID),
			Name: band.Name + " - " + band.Country.Name + " (" + strconv.Itoa(band.FormedIn) + ")",
		})
	}

	return collection
}
