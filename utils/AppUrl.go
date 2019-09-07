package utils

import (
	"net/url"
	"os"
)

var AppUrl = url.URL{Scheme: "http", Host: os.Getenv("APP_HOST")}

func MakeAppUrl(path string) string {
	var localUrl = AppUrl
	localUrl.Path = path
	return localUrl.String()
}
