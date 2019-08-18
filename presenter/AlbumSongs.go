package presenter

import (
	"github.com/vortgo/ma-parser/models"
	"sort"
)

type songInCollection struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Time     string `json:"time"`
	Lyrics   string `json:"lyrics"`
	Position int    `json:"position"`
}

type albumSongsPresenter struct {
	songs []*models.Song
}

func NewAlbumSongsPresenter(songs []*models.Song) *albumSongsPresenter {
	return &albumSongsPresenter{songs: songs}
}

func (presenter *albumSongsPresenter) SongsCollection() []*songInCollection {
	var collection []*songInCollection
	for _, song := range presenter.songs {
		collection = append(collection, &songInCollection{
			ID:       int(song.ID),
			Name:     song.Name,
			Time:     song.Time,
			Lyrics:   song.Lyrics,
			Position: song.Position,
		})
	}

	sort.Slice(collection, func(i, j int) bool {
		return collection[i].Position < collection[j].Position
	})
	return collection
}
