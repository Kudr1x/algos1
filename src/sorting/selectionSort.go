package sorting

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minValue := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minValue] {
				minValue = j
			}
		}

		arr[i], arr[minValue] = arr[minValue], arr[i]
	}

	return arr
}
