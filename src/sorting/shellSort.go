package sorting

func ShellSort(arr []int) []int {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j := i; j >= gap && arr[j-gap] > arr[j]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}

	return arr
}

func ShellSortHibbard(arr []int) []int {
	var gaps []int
	g := 1
	for g < len(arr) {
		gaps = append(gaps, g)
		g = g*2 + 1
	}

	for k := len(gaps) - 1; k >= 0; k -= 1 {
		gap := gaps[k]
		for i := gap; i < len(arr); i++ {
			for j := i; j >= gap && arr[j-gap] > arr[j]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}

	return arr
}

func ShellSortPratt(arr []int) []int {
	var gaps []int
	for h := 1; h <= len(arr); h *= 2 {
		for k := 0; k <= len(arr)/h; k++ {
			gap := h * (1 << k)
			if gap > len(arr) {
				break
			}
			gaps = append(gaps, gap)
		}
	}

	for k := len(gaps) - 1; k >= 0; k -= 1 {
		gap := gaps[k]
		for i := gap; i < len(arr); i++ {
			for j := i; j >= gap && arr[j-gap] > arr[j]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}

	return arr
}
