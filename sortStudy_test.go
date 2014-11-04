package sortStudy

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestIsSorted(t *testing.T) {
	data := ints
	dataSlice := IntSlice(data[0:])
	if IsSorted(dataSlice) {
		t.Error("Array %v was marked as sorted!", ints)
	}
}

func TestInsertionSort(t *testing.T) {
	data := ints
	dataSlice := IntSlice(data[0:])
	InsertionSort(dataSlice)
	fmt.Println(dataSlice)
	if !IsSorted(dataSlice) {
		t.Error("Input %v", ints)
		t.Error("Sorted %v", data)
	}
}

func TestSelectionSort(t *testing.T) {
	data := ints
	dataSlice := IntSlice(data[0:])
	SelectionSort(dataSlice)
	fmt.Println(dataSlice)
	if !IsSorted(dataSlice) {
		t.Error("Input %v", ints)
		t.Error("Sorted %v", data)
	}
}

func TestMergeSort(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	size := 3000000
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	//	data := ints
	dataSlice := IntSlice(data[0:])
	buffer := make([]int, len(data))
	MergeSort(dataSlice, IntSlice(buffer))
	//	fmt.Println(dataSlice)
	if !IsSorted(dataSlice) {
		t.Error("Input %v", ints)
		t.Error("Sorted %v", data)
	}
}
