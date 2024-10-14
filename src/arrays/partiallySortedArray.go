package arrays

func PartiallySortedArray(length int) []int {
	var arr []int

	arr = SortedArray(length)

	for i, j := (length/10)*9, length-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
