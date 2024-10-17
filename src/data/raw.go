package data

import (
	"Algos1/src/arrays"
	"Algos1/src/sorting"
)

var DSort = []SortData{
	{Name: "BubbleSort", Fun: sorting.BubbleSort},
	{Name: "HeapSort", Fun: sorting.HeapSort},
	{Name: "InsertionSort", Fun: sorting.InsertionSort},
	{Name: "MergeSort", Fun: sorting.MergeSort},
	{Name: "QuickSort", Fun: sorting.QuickSort},
	{Name: "SelectionSort", Fun: sorting.SelectionSort},
	{Name: "ShellSort", Fun: sorting.ShellSort},
	{Name: "ShellSortHibbard", Fun: sorting.ShellSortHibbard},
	{Name: "ShellSortPratt", Fun: sorting.ShellSortPratt},
}

var DArray = []ArrayData{
	{Name: "Случайный массив", Fun: arrays.RandomArray},
	{Name: "Отсортированный массив", Fun: arrays.SortedArray},
	{Name: "Обратно отсортированный массив", Fun: arrays.ReverseSortedArray},
	{Name: "Частично отсортированный массив", Fun: arrays.PartiallySortedArray},
}
