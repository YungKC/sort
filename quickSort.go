// Package sort to provide known types of sorting algorithm

package sortStudy

import "fmt"

//import "reflect"
import "math/rand"

// MergeSort
func QuickSort(data Interface) {
	quickSort(data, 0, data.Len()-1)
}

// merge sort data from index a to b

func quickSort(data Interface, start, end int) {
	fmt.Println("quickSort: ", data, start, end)
	if start == end || end < 0 || start < 0 {
		return
	} else if end-start == 1 {
		if data.Less(end, start) {
			data.Swap(start, end)
		}
		return
	}

	pivit := start + rand.Intn(end-start)
	fmt.Println("pivit: ", pivit)

	data.Swap(pivit, end)
	fmt.Println("swapped: ", data)

	pivit = start
	for i := start + 1; i < end-1; i++ {
		if data.Less(i, end) {
			data.Swap(i, pivit+1)
			pivit = i
			fmt.Println("  Swapped: ", i, data.Get(i), data)

		}
	}

	fmt.Println("  before last Swap: ", pivit, data)
	data.Swap(end, pivit)
	fmt.Println("  pivited: ", pivit, data)
	quickSort(data, 0, pivit-1)
	quickSort(data, pivit+1, end)

}
