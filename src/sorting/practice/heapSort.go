package practice

func HeapSort(arr []int) []int {
	len := len(arr)
	half := len / 2

	for i := half; i > -1; i-- {
		heap(arr, i, len-1)
	}

	for j := len - 1; j > 0; j-- {
		arr[j], arr[0] = arr[0], arr[j]
		heap(arr, 0, j-1)
	}
	return arr
}

func heap(array []int, i, end int) {
	left := 2*i + 1
	if left > end {
		return
	}
	tmp := left
	right := 2*i + 2

	if right <= end && array[tmp] < array[right] {
		tmp = right
	}

	if array[i] < array[tmp] {
		array[i], array[tmp] = array[tmp], array[i]
		heap(array, tmp, end)
	}
}
