package sortStudy

import (
	"fmt"
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
	data := ints
	dataSlice := IntSlice(data[0:])
	MergeSort(dataSlice)
	fmt.Println(dataSlice)
	if !IsSorted(dataSlice) {
		t.Error("Input %v", ints)
		t.Error("Sorted %v", data)
	}
}
