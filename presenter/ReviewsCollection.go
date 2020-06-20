package presenter

import (
	"github.com/vortgo/ma-parser/models"
	"sort"
)

type reviewInCollection struct {
	ID         int     `json:"id"`
	Title      *string `json:"title"`
	Author     *string `json:"author"`
	Rating     *int    `json:"rating"`
	AlbumName  *string `json:"album_name,omitempty"`
	AlbumImage *string `json:"album_image,omitempty"`
	BandName   *string `json:"band_name,omitempty"`
}

type reviewsCollectionPresenter struct {
	reviews []*models.Review
}

func NewReviewsCollectionPresenter(reviews []*models.Review) *reviewsCollectionPresenter {
	return &reviewsCollectionPresenter{reviews: reviews}
}

func (p *reviewsCollectionPresenter) BandReviewsCollection() []reviewInCollection {
	collection := make([]reviewInCollection, 0)
	for _, review := range p.reviews {
		collection = append(collection, reviewInCollection{
			ID:        int(review.ID),
			AlbumName: &review.Album.Name,
			Title:     &review.Title,
			Author:    &review.Author,
			Rating:    &review.Rating,
		})
	}

	sort.Slice(collection, func(i, j int) bool {
		return *collection[i].Rating > *collection[j].Rating
	})
	return collection
}

func (p *reviewsCollectionPresenter) AlbumReviewsCollection() []reviewInCollection {
	collection := make([]reviewInCollection, 0)
	for _, review := range p.reviews {
		collection = append(collection, reviewInCollection{
			ID:     int(review.ID),
			Title:  &review.Title,
			Author: &review.Author,
			Rating: &review.Rating,
		})
	}

	sort.Slice(collection, func(i, j int) bool {
		return *collection[i].Rating > *collection[j].Rating
	})
	return collection
}

func (p *reviewsCollectionPresenter) ReviewsCollection() []reviewInCollection {
	collection := make([]reviewInCollection, 0)
	for _, review := range p.reviews {
		collection = append(collection, reviewInCollection{
			ID:         int(review.ID),
			Title:      &review.Title,
			Author:     &review.Author,
			Rating:     &review.Rating,
			AlbumName:  &review.Album.Name,
			AlbumImage: &review.Album.Image,
			BandName:   &review.Album.Band.Name,
		})
	}

	sort.Slice(collection, func(i, j int) bool {
		return *collection[i].Rating > *collection[j].Rating
	})
	return collection
}
