package util

type SortData struct {
	Name string
	Fun  func(arr []int) []int
}

type ArrayData struct {
	Name string
	Fun  func(length int) []int
}
