package theory

import (
	"math"
)

func BestQuick(n int) int {
	return n * int(math.Log(float64(n)))
}

func AverageQuick(n int) int {
	return int(math.Log2(float64(n))) * (n + 1)
}

func WurstQuick(n int) int {
	return (n*n + n) / 2
}
