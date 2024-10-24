package theory

import (
	"math"
)

func BestMerge(n int) int {
	return n*int(math.Log2(float64(n))) + n
}

func AverageMerge(n int) int {
	return n*int(math.Log2(float64(n))) + n
}

func WurstMerge(n int) int {
	return n*int(math.Log2(float64(n))) + n
}
