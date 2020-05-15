package handler

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"ma-api/presenter"
	"net/http"
	"strconv"
)

func Review(c echo.Context) error {
	var review models.Review
	reviewId, _ := strconv.Atoi(c.Param("id"))

	db := repositories.PostgresDB.Find(&review, reviewId)

	if db.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	simpleReview := presenter.NewReviewPresenter(&review).SimpleReview()

	return c.JSON(200, &simpleReview)
}
