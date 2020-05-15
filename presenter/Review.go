package presenter

import (
	"github.com/vortgo/ma-parser/models"
)

type review struct {
	ID      uint    `json:"id"`
	AlbumId uint    `json:"album_id"`
	Title   *string `json:"title"`
	Author  *string `json:"author"`
	Rating  *int    `json:"rating"`
}

type reviewPresenter struct {
	review *models.Review
}

func NewReviewPresenter(review *models.Review) *reviewPresenter {
	return &reviewPresenter{review: review}
}

func (p *reviewPresenter) SimpleReview() review {
	return review{
		ID:      p.review.ID,
		AlbumId: p.review.AlbumID,
		Title:   &p.review.Title,
		Author:  &p.review.Author,
		Rating:  &p.review.Rating,
	}
}
