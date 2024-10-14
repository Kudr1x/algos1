package plot

import (
	"Algos1/src/util"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func DrawLine(sortData util.SortData) {
	p := plot.New()

	p.Title.Text = sortData.Name
	p.X.Label.Text = "N"
	p.Y.Label.Text = "Мс"
	p.X.Label.Position = 1
	p.Y.Label.Position = 1

	p.Y.Min = 0
	p.Y.Max = 80
	p.X.Min = 0
	p.X.Max = 200000

	for _, arrData := range util.DArray {
		line := plotter.XYs{}

		for i := 1000; i <= 100000; i += 10000 {
			line = append(line, plotter.XY{X: float64(i), Y: float64(util.MeasureTime(sortData.Fun, arrData.Fun, i))})
		}

		p.Add(plotter.NewGrid())
		_ = plotutil.AddLinePoints(p, arrData.Name, line)
	}

	DisplayPlot(p, sortData.Name)
}
