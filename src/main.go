package main

import (
	"Algos1/src/plot"
	"Algos1/src/util"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, i := range util.DSort {
		wg.Add(1)

		go func(i util.SortData) {
			defer wg.Done()
			plot.DrawLine(i)
		}(i)
	}

	wg.Wait()
}
