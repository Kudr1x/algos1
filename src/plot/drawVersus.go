package plot

import (
	"Algos1/src/data"
	"os/exec"
	"strconv"
	"strings"
)

func DrawVersus() {
	drawTheoryVersus()
}

func drawTheoryVersus() {
	var dataXY []string

	for _, s := range data.DSortTheory {
		var arrX []string
		var arrY []string

		for i := 1000; i <= 101000; i += 10000 {
			tempX := strconv.Itoa(i)
			tempY := strconv.Itoa(s.Average(i))
			arrX = append(arrX, tempX)
			arrY = append(arrY, tempY)
		}

		dataX := strings.Join(arrX, ",")
		dataY := strings.Join(arrY, ",")

		dataXY = append(dataXY, dataX)
		dataXY = append(dataXY, dataY)
	}

	data := strings.Join(dataXY, ";")
	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos1/src/plot/displayVersusPlot.py", data)
	cmd.CombinedOutput()
}
