package arrays

import (
	"math/rand"
)

func RandomArray(length int) []int {
	var arr []int

	rand.Seed(51)

	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(length))
	}

	return arr
}
