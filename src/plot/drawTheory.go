package plot

import (
	"Algos1/src/data"
	"os/exec"
	"strconv"
	"strings"
)

func DrawTheory(sorts data.TheoryData) {
	var dataXY []string
	var name = sorts.Name

	for t := 0; t < 3; t++ {
		var arrX []string
		var arrY []string

		for i := 1000; i <= 101000; i += 10000 {
			tempX := strconv.Itoa(i)
			tempY := ""
			if t == 0 {
				tempY = strconv.Itoa(sorts.Best(i))
			} else if t == 1 {
				tempY = strconv.Itoa(sorts.Average(i))
			} else if t == 2 {
				tempY = strconv.Itoa(sorts.Wurst(i))
			}
			arrX = append(arrX, tempX)
			arrY = append(arrY, tempY)
		}

		dataX := strings.Join(arrX, ",")
		dataY := strings.Join(arrY, ",")

		dataXY = append(dataXY, dataX)
		dataXY = append(dataXY, dataY)
	}

	data := strings.Join(dataXY, ";")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos1/src/plot/displayTheoryPlot.py", data, name)

	cmd.CombinedOutput()
}
