package arrays

import (
	"sort"
)

func SortedArray(length int) []int {
	var arr []int

	arr = RandomArray(length)

	sort.Ints(arr)

	return arr
}
