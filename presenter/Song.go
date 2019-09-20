package presenter

import (
	"github.com/vortgo/ma-parser/models"
)

type songPresenter struct {
	song *models.Song
}

type simpleSong struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Time      string  `json:"time"`
	Lyrics    string  `json:"lyrics"`
	Position  int     `json:"position"`
	AlbumName string  `json:"album_name"`
	AlbumId   int     `json:"album_id"`
	BandName  string  `json:"band_name"`
	BandId    int     `json:"band_id"`
	VideoKey  *string `json:"video_key"`
}

func NewSongPresenter(song *models.Song) *songPresenter {
	return &songPresenter{song: song}
}

func (presenter *songPresenter) SimpleSongs(videoKey string) *simpleSong {
	var videoKeyPointer *string
	if videoKey == "" {
		videoKeyPointer = nil
	} else {
		videoKeyPointer = &videoKey
	}

	return &simpleSong{
		Id:        int(presenter.song.ID),
		Name:      presenter.song.Name,
		Time:      presenter.song.Time,
		Lyrics:    presenter.song.Lyrics,
		Position:  presenter.song.Position,
		AlbumName: presenter.song.Album.Name,
		AlbumId:   int(presenter.song.AlbumID),
		BandName:  presenter.song.Band.Name,
		BandId:    int(presenter.song.BandID),
		VideoKey:  videoKeyPointer,
	}
}
