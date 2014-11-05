// Package sort to provide known types of sorting algorithm

package sort

//import "fmt"
//import "reflect"

// MergeSort
// TODO: find out how to allocate buffer with Interface type instead of IntSlice
func MergeSort(data Interface) {
	buffer := Interface(make(IntSlice, data.Len()))
	mergeSort(data, buffer, 0, data.Len()-1, false)
}

func MergeSortConcurrent(data Interface) {
	buffer := Interface(make(IntSlice, data.Len()))
	mergeSort(data, buffer, 0, data.Len()-1, true)
}

func mergeSort(data, buffer Interface, a, b int, isConcurrent bool) {
	//	fmt.Println("mergeSort: ", a, b, data)
	if a == b {
		return
	} else if b-a == 1 {
		if data.Less(b, a) {
			data.Swap(a, b)
		}
		return
	}

	center := (a + b) / 2

	if isConcurrent {
		// parallelize the split merge logic here
		c := make(chan bool)
		go func() { mergeSort(data, buffer, a, center, isConcurrent); c <- true }()
		go func() { mergeSort(data, buffer, center+1, b, isConcurrent); c <- true }()
		for i := 0; i < 2; i++ {
			<-c
		}
		close(c)
	} else {
		// serial implementation
		mergeSort(data, buffer, a, center, isConcurrent)
		mergeSort(data, buffer, center+1, b, isConcurrent)
	}
	merge(data, buffer, a, center, b)
}

func merge(data, buffer Interface, left, center, right int) {

	//	fmt.Println("merge: ", left, center, right, data)

	leftIndex := left
	rightIndex := center + 1
	bufferIndex := left
	leftReached := false
	rightReached := false
	for bufferIndex <= right {
		if data.Less(leftIndex, rightIndex) && !leftReached || rightReached {
			buffer.Set(bufferIndex, data.Get(leftIndex))
			bufferIndex++
			if leftIndex < center {
				leftIndex++
			} else {
				leftReached = true
			}
		} else {
			buffer.Set(bufferIndex, data.Get(rightIndex))
			bufferIndex++
			if rightIndex < right {
				rightIndex++
			} else {
				rightReached = true
			}
		}
	}
	//	fmt.Println("Buffer: ", buffer)
	bufferIndex = left
	for i := left; i <= right; i++ {
		data.Set(i, buffer.Get(bufferIndex))
		bufferIndex++
	}

}
