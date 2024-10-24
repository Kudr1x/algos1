package main

import (
	"Algos1/src/data"
	"Algos1/src/plot"
	"sync"
)

func launch() {
	//practiceGraphs()
	//theoryGraphs()
	versusGraphs()
}

func versusGraphs() {
	plot.DrawVersus()
}

func practiceGraphs() {
	var wg sync.WaitGroup

	for _, i := range data.DSort {
		wg.Add(1)

		go func(i data.SortData) {
			defer wg.Done()
			plot.DrawPractice(i)
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
			plot.DrawTheory(i)
		}(i)
	}

	wg.Wait()
}

func main() {
	launch()
}
