package theory

func BestBubble(n int) int {
	return n * (n - 1) / 4
}

func AverageBubble(n int) int {
	return n * (n + 1) / 2
}

func WurstBubble(n int) int {
	return n * (n + 1)
}
