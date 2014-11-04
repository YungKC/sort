// Package sort to provide known types of sorting algorithm

package sortStudy

//import "fmt"

// MergeSort
func MergeSort(data []int) {
	buffer := make([]int, len(data))
	mergeSort(data, buffer, 0, len(data)-1)
}

// merge sort data from index a to b

func mergeSort(data, buffer []int, a, b int) {
	//	fmt.Println("mergeSort: ", a, b, data)
	if a == b {
		return
	} else if b-a == 1 {
		if data[a] > data[b] {
			tmp := data[a]
			data[a] = data[b]
			data[b] = tmp
		}
		return
	}

	center := (a + b) / 2
	mergeSort(data, buffer, a, center)
	mergeSort(data, buffer, center+1, b)
	merge(data, buffer, a, center, b)
}

func merge(data, buffer []int, left, center, right int) {

	//	fmt.Println("merge: ", left, center, right, data)

	leftIndex := left
	rightIndex := center + 1
	bufferIndex := left
	leftReached := false
	rightReached := false
	for bufferIndex <= right {
		if data[leftIndex] < data[rightIndex] && !leftReached || rightReached {
			buffer[bufferIndex] = data[leftIndex]
			bufferIndex++
			if leftIndex < center {
				leftIndex++
			} else {
				leftReached = true
			}
		} else {
			buffer[bufferIndex] = data[rightIndex]
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
		data[i] = buffer[bufferIndex]
		bufferIndex++
	}

}
