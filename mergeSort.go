// Package sort to provide known types of sorting algorithm

package sortStudy

//import "fmt"
//import "reflect"

// MergeSort
func MergeSort(data, buffer Interface) {
	mergeSort(data, buffer, 0, data.Len()-1)
}

// merge sort data from index a to b

func mergeSort(data, buffer Interface, a, b int) {
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

	// parallelize the split merge logic here
	c := make(chan bool)
	go func() { mergeSort(data, buffer, a, center); c <- true }()
	go func() { mergeSort(data, buffer, center+1, b); c <- true }()
	for i := 0; i < 2; i++ {
		<-c
	}

	// serial implementation
	//mergeSort(data, buffer, a, center)
	//mergeSort(data, buffer, center+1, b)

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
