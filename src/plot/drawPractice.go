package plot

import (
	"Algos1/src/data"
	"Algos1/src/util"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func DrawLine(sortData data.SortData) {
	var dataXY []string

	for _, arrData := range data.DArray {
		var arrX []string
		var arrY []string

		for i := 1000; i <= 101000; i += 10000 {
			tempX := strconv.Itoa(i)
			tempY := strconv.FormatInt(util.MeasureTime(sortData.Fun, arrData.Fun, i), 10)
			arrX = append(arrX, tempX)
			arrY = append(arrY, tempY)

			fmt.Println(sortData.Name, tempX, tempY)
		}

		dataX := strings.Join(arrX, ",")
		dataY := strings.Join(arrY, ",")

		dataXY = append(dataXY, dataX)
		dataXY = append(dataXY, dataY)
	}

	data := strings.Join(dataXY, ";")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos1/src/plot/displayPlotWithRegression.py", data, sortData.Name)

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(out)
	}
}
