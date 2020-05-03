package handler

import (
	"github.com/labstack/echo"
	"github.com/vortgo/ma-parser/models"
	"github.com/vortgo/ma-parser/repositories"
	"time"
)

func ApiStat(c echo.Context) error {
	db := repositories.PostgresDB
	now := time.Now().Truncate(24 * time.Hour)
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	tomorrow := today.AddDate(0, 0, 1)
	urls, _ := db.Table("api_requests").Select("url, count(url) as count").Where("created_at > ? AND created_at < ?", today, tomorrow).Group("url").Rows()

	var urlStat = make(map[string]int)

	for urls.Next() {
		var url string
		var count int
		urls.Scan(&url, &count)
		urlStat[url] = count
	}

	var totalRequests int
	db.Table("api_requests").Where("created_at > ? AND created_at < ?", today, tomorrow).Count(&totalRequests)

	var uniqueIpCount int
	db.Table("api_requests").Where("created_at > ? AND created_at < ?", today, tomorrow).Group("ip").Count(&uniqueIpCount)

	var latestBandUpdate models.LatestBandUpdate
	db.Order("updated_at desc").First(&latestBandUpdate)
	var upcomingAlbum models.UpcomingAlbum
	db.Order("updated_at desc").First(&upcomingAlbum)

	var response = make(map[string]interface{})
	response["urls"] = urlStat
	response["total_requests"] = totalRequests
	response["uniq_by_ip"] = uniqueIpCount
	response["last_band_upd_updated_at"] = latestBandUpdate.UpdatedAt.Format("2006.01.02-15.04.05")
	response["last_upcoming_album_updated_at"] = upcomingAlbum.UpdatedAt.Format("2006.01.02-15.04.05")

	return c.JSON(200, response)
}
