// Package sort to provide known types of sorting algorithm

package sortStudy

//import "fmt"

// MergeSort
func MergeSort(data []int) {
	mergeSort(data, 0, len(data)-1)
}

// merge sort data from index a to b

func mergeSort(data []int, a, b int) {
	//	fmt.Println("mergeSort: ", a, b, data)
	if a == b {
		return
	}

	center := (a + b) / 2
	mergeSort(data, a, center)
	mergeSort(data, center+1, b)
	merge(data, a, center, b)
}

func merge(data []int, left, center, right int) {

	//	fmt.Println("merge: ", left, center, right, data)

	if right-left == 1 {
		if data[left] > data[right] {
			tmp := data[left]
			data[left] = data[right]
			data[right] = tmp
		}
		return
	}

	buffer := make([]int, right-left+1)
	leftIndex := left
	rightIndex := center + 1
	bufferIndex := 0
	leftReached := false
	rightReached := false
	for bufferIndex < len(buffer) {
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
	bufferIndex = 0
	for i := left; i <= right; i++ {
		data[i] = buffer[bufferIndex]
		bufferIndex++
	}

}
