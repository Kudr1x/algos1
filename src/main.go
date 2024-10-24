package main

import (
	"Algos1/src/data"
	"Algos1/src/plot"
	"sync"
)

func main() {
	launch()
}

func launch() {
	//practiceGraphs()
	theoryGraphs()
}

func practiceGraphs() {
	var wg sync.WaitGroup

	for _, i := range data.DSort {
		wg.Add(1)

		go func(i data.SortData) {
			defer wg.Done()
			plot.DrawLine(i)
		}(i)
	}

	wg.Wait()
}

func theoryGraphs() {
	var wg sync.WaitGroup

	for _, i := range data.DSortTheory {
		wg.Add(1)

		go func(i data.TheoryData) {
			defer wg.Done()
			plot.DrawLineTheory(i)
		}(i)
	}

	wg.Wait()
}
