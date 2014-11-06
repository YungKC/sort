// Package sort to provide known types of sorting algorithm
package sort

//import "reflect"

// MergeSort
func HeapSort(data Interface) {
	size := data.Len()
	// get last parent node
	for parentNode := (size + 2) / 2; parentNode >= 0; parentNode-- {
		sink(data, parentNode, size)
	}

	// enforce data[0:end] is in heap order
	// sorted array is from [end: count]
	end := size - 1
	for {
		if end <= 0 {
			return
		}
		data.Swap(0, end)
		end--
		sink(data, 0, end+1)
	}
}

func sink(data Interface, index, size int) {
	nextLevel := 2 * index
	if nextLevel+2 > size {
		return
	}
	if data.Less(index, nextLevel+1) || (nextLevel+2 < size && data.Less(index, nextLevel+2)) {
		if nextLevel+2 >= size || data.Less(nextLevel+2, nextLevel+1) {
			data.Swap(index, 2*index+1)
			sink(data, 2*index+1, size)
		} else {
			data.Swap(index, 2*index+2)
			sink(data, 2*index+2, size)
		}
	}
}
