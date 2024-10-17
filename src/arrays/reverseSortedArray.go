package arrays

func ReverseSortedArray(length int) []int {
	var arr []int

	arr = SortedArray(length)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}