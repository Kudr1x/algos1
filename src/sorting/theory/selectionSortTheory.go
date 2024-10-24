package theory

func BestSelection(n int) int {
	return n * (n - 1) / 2
}

func AverageSelection(n int) int {
	return n * (n + 3) / 4
}

func WurstSelection(n int) int {
	return n * (n + 1) / 2
}
