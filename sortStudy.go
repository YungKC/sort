// Package sort to provide known types of sorting algorithm
package sortStudy

// Data interface that supports sorting
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func InsertionSort(data Interface) {
	insertionSort(data, 0, data.Len())
}

// sort data from index a to b
func insertionSort(data Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

func IsSorted(data Interface) bool {
	length := data.Len()
	for i := 1; i < length; i++ {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
