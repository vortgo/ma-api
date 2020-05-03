package middleware

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/repositories"
	"time"
)

type ApiRequests struct {
	ID        uint
	URL       string
	Ip        string
	UserAgent string
	CreatedAt time.Time
}

func RequestStatisticMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiRequest := ApiRequests{
			URL:       c.Request().URL.String(),
			Ip:        c.RealIP(),
			UserAgent: c.Request().UserAgent(),
		}

		repositories.PostgresDB.Save(&apiRequest)

		return nil
	}
}
