package sortStudy

import (
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestIntsInsertionSort(t *testing.T) {
	data := ints
	dataSlice := IntSlice(data[0:])
	InsertionSort(dataSlice)
	if !IsSorted(dataSlice) {
		t.Error("Sorted %v", ints)
		t.Error("Sorted %v", data)
	}
}
