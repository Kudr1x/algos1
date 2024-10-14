package util

import (
	"time"
)

func MeasureTime(sortFun func(arr []int) []int, arrFun func(length int) []int, length int) int64 {
	var arr = arrFun(length)

	start := time.Now()

	sortFun(arr)

	return time.Since(start).Milliseconds()
}
