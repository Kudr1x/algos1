package theory

import (
	"math"
)

func BestInsertion(n int) int {
	return 4 * (n - 1)
}

func AverageInsertion(n int) int {
	return ((n * n) + int(math.Log(float64(n)))) / 2
}

func WurstInsertion(n int) int {
	return (n*n)/4 + int(math.Log(float64(n)))
}
