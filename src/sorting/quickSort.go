package sorting

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var less []int
	var greater []int

	for i, v := range arr {
		if i != pivotIndex {
			if v < pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
	}

	sortedLess := QuickSort(less)
	sortedGreater := QuickSort(greater)
	sorted := append(sortedLess, pivot)
	sorted = append(sorted, sortedGreater...)

	return sorted
}
