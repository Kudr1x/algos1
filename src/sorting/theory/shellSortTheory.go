package theory

import (
	"math"
)

func BestShell(n int) int {
	return n * int(math.Log(float64(n)))
}

func AverageShell(n int) int {
	return int(math.Pow(float64(n), 3/2))
}

func WurstShell(n int) int {
	return n * n
}

func BestShellPratt(n int) int {
	return n * int(math.Log(float64(n)))
}

func AverageShellPratt(n int) int {
	return n * int(math.Pow(math.Log(float64(n)), 2))
}

func WurstShellPratt(n int) int {
	return n * int(math.Pow(math.Log(float64(n)), 2))
}

func BestShellHibbard(n int) int {
	return n * int(math.Log(float64(n)))
}

func AverageShellHibbard(n int) int {
	return int(math.Pow(float64(n), 5/4))
}

func WurstShellHibbard(n int) int {
	return int(math.Pow(float64(n), 3/2))
}
