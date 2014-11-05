package sort_test

import (
	"fmt"
	. "github.com/yungkc/sort"
	"math"
	"math/rand"
	"runtime"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, math.MaxInt64, math.MinInt64, 7586, -5467984, 7586}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Set GOMAXPROCS to ", runtime.NumCPU())
}

func TestIsSorted(t *testing.T) {
	data := ints
	dataSlice := IntSlice(data[0:])
	if IsSorted(dataSlice) {
		t.Error("Array %v was marked as sorted!", ints)
	}
}

func testInternal(t *testing.T, searchFn SortFn) {
	data := ints
	dataSlice := IntSlice(data[0:])
	searchFn(dataSlice)
	//	fmt.Println(dataSlice)
	if !IsSorted(dataSlice) {
		t.Error("Input %v", ints)
		t.Error("Sorted %v", data)
	}
}

func TestInsertionSort(t *testing.T) {
	testInternal(t, InsertionSort)
}

func TestSelectionSort(t *testing.T) {
	testInternal(t, SelectionSort)
}

func TestQuickSort(t *testing.T) {
	testInternal(t, QuickSort)
}

func TestMergeSort(t *testing.T) {
	testInternal(t, MergeSort)
}

func TestMergeSortConcurrent(t *testing.T) {
	testInternal(t, MergeSortConcurrent)
}

func generateRandomArray(count int) []int {
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Int()
	}
	return data
}

func benchmarkInternal(b *testing.B, searchFn SortFn, size int) {
	data := generateRandomArray(size)
	dataSlice := IntSlice(data[0:])
	b.ResetTimer()
	searchFn(dataSlice)
	b.StopTimer()
	if !IsSorted(dataSlice) {
		b.Error("Input %v", ints)
		b.Error("Sorted %v", data)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkInternal(b, MergeSort, 2000000)
}

func BenchmarkMergeSortConcurrent(b *testing.B) {
	benchmarkInternal(b, MergeSortConcurrent, 2000000)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarkInternal(b, QuickSort, 2000000)
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkInternal(b, InsertionSort, 10000)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkInternal(b, SelectionSort, 10000)
}
