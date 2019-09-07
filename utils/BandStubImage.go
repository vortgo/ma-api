package utils

import (
	"math"
	"strconv"
)

func GetBandStabImageUrl(id int) string {
	numberImage := int(math.Mod(float64(id), 4))
	path := "static/empty_band/" + strconv.Itoa(numberImage) + ".jpg"
	return MakeAppUrl(path)
}
