package theory

import "math"

func BestHeap(n int) int {
	return n
}

func AverageHeap(n int) int {
	return n * int(math.Log(float64(n)))
}

func WurstHeap(n int) int {
	return n * int(math.Log(float64(n)))
}
