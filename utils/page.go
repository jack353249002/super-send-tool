package utils

import (
	"math"
)

func CreatePage(total, pageSize int) (totalPage int) {
	totalPageTemp := math.Ceil(float64(total) / float64(pageSize))
	totalPage = int(totalPageTemp)
	return
}
func CreatePageOffset(page, pageSize int) (offset int) {
	offset = (page - 1) * pageSize
	return
}
