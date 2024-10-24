package data

type SortData struct {
	Name string
	Fun  func(arr []int) []int
}

type ArrayData struct {
	Name string
	Fun  func(length int) []int
}

type TheoryData struct {
	Name    string
	Best    func(n int) int
	Average func(n int) int
	Wurst   func(n int) int
}

type TheorySorts struct {
	Sorts TheoryData
}
