package presenter

import (
	"github.com/vortgo/ma-parser/models"
)

type collectionSongPresenter struct {
	songs []*models.Song
}

type searchSong struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	BandName string `json:"band_name"`
}

func NewCollectionSongPresenter(songs []*models.Song) *collectionSongPresenter {
	return &collectionSongPresenter{songs: songs}
}

func (presenter *collectionSongPresenter) SearchSongs() []*searchSong {
	var collection []*searchSong

	for _, song := range presenter.songs {
		collection = append(collection, &searchSong{
			ID:       int(song.ID),
			Name:     song.Name,
			BandName: song.Band.Name,
		})
	}

	return collection
}
