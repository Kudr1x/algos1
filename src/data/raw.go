package data

import (
	"Algos1/src/arrays"
	"Algos1/src/sorting/practice"
	"Algos1/src/sorting/theory"
)

var DSort = []SortData{
	{Name: "BubbleSort", Fun: practice.BubbleSort},
	{Name: "HeapSort", Fun: practice.HeapSort},
	{Name: "InsertionSort", Fun: practice.InsertionSort},
	{Name: "MergeSort", Fun: practice.MergeSort},
	{Name: "QuickSort", Fun: practice.QuickSort},
	{Name: "SelectionSort", Fun: practice.SelectionSort},
	{Name: "ShellSort", Fun: practice.ShellSort},
	{Name: "ShellSortHibbard", Fun: practice.ShellSortHibbard},
	{Name: "ShellSortPratt", Fun: practice.ShellSortPratt},
}

var DArray = []ArrayData{
	{Name: "Случайный массив", Fun: arrays.RandomArray},
	{Name: "Отсортированный массив", Fun: arrays.SortedArray},
	{Name: "Обратно отсортированный массив", Fun: arrays.ReverseSortedArray},
	{Name: "Частично отсортированный массив", Fun: arrays.PartiallySortedArray},
}
var bubbleSortTheory = TheoryData{
	Name:    "BubbleSort",
	Best:    theory.BestBubble,
	Average: theory.AverageBubble,
	Wurst:   theory.WurstBubble,
}

var insertionSortTheory = TheoryData{
	Name:    "InsertionSort",
	Best:    theory.BestInsertion,
	Average: theory.AverageInsertion,
	Wurst:   theory.WurstInsertion,
}

var mergeSortTheory = TheoryData{
	Name:    "MergeSort",
	Best:    theory.BestMerge,
	Average: theory.AverageMerge,
	Wurst:   theory.WurstMerge,
}

var quickSortTheory = TheoryData{
	Name:    "QuickSort",
	Best:    theory.BestQuick,
	Average: theory.AverageQuick,
	Wurst:   theory.WurstQuick,
}

var selectionSortTheory = TheoryData{
	Name:    "SelectionSort",
	Best:    theory.BestSelection,
	Average: theory.AverageSelection,
	Wurst:   theory.WurstSelection,
}

var shellSortTheory = TheoryData{
	Name:    "ShellSort",
	Best:    theory.BestShell,
	Average: theory.AverageShell,
	Wurst:   theory.WurstShell,
}

var shellSortPrattTheory = TheoryData{
	Name:    "ShellSortPratt",
	Best:    theory.BestShellPratt,
	Average: theory.AverageShellPratt,
	Wurst:   theory.WurstShellPratt,
}

var shellSortHibbardTheory = TheoryData{
	Name:    "ShellSortHibbard",
	Best:    theory.BestShellHibbard,
	Average: theory.AverageShellHibbard,
	Wurst:   theory.WurstShellHibbard,
}

var heapSortTheory = TheoryData{
	Name:    "HeapSort",
	Best:    theory.BestHeap,
	Average: theory.AverageHeap,
	Wurst:   theory.WurstHeap,
}

var DSortTheory = []TheoryData{
	selectionSortTheory,
	bubbleSortTheory,
	insertionSortTheory,
	mergeSortTheory,
	quickSortTheory,
	shellSortTheory,
	shellSortPrattTheory,
	shellSortHibbardTheory,
	heapSortTheory,
}
