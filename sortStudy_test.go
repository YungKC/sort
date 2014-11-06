package sort_test

import (
	"fmt"
	. "github.com/yungkc/sort"
	"math"
	"math/rand"
	"runtime"
	"testing"
)

var ints = [...]int{99999, 74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, math.MaxInt64, math.MinInt64, 7586, -5467984, 7586}
var randomInts []int
var benchmarkSize = 100000

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Set GOMAXPROCS to ", runtime.NumCPU())
	randomInts = generateRandomArray(benchmarkSize)
}

func TestIsSorted(t *testing.T) {
	data := ints
	dataSlice := IntSlice(data[0:])
	if IsSorted(dataSlice) {
		t.Error("Array %v was marked as sorted!", ints)
	}
}

func testInternal(t *testing.T, searchFn SortFn) {
	// data := ints
	data := generateRandomArray(10000)
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

func TestBubbleSort(t *testing.T) {
	testInternal(t, BubbleSort)
}

func TestHeapSort(t *testing.T) {
	testInternal(t, HeapSort)
}

func generateRandomArray(count int) []int {
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = rand.Intn(count) - count/2
	}
	return data
}

func benchmarkInternal(b *testing.B, searchFn SortFn, size int) {
	for n := 0; n < b.N; n++ {
		//	data := generateRandomArray(size)
		data := make([]int, size)
		for i := 0; i < size; i++ {
			data[i] = randomInts[i]
		}
		dataSlice := IntSlice(data[0:])
		//	fmt.Println(dataSlice)
		b.StartTimer()
		searchFn(dataSlice)
		//	fmt.Println(dataSlice)

		b.StopTimer()
		if !IsSorted(dataSlice) {
			b.Error("Input %v", ints)
			b.Error("Sorted %v", data)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	benchmarkInternal(b, InsertionSort, benchmarkSize)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkInternal(b, SelectionSort, benchmarkSize)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkInternal(b, MergeSort, benchmarkSize)
}

func BenchmarkMergeSortConcurrent(b *testing.B) {
	benchmarkInternal(b, MergeSortConcurrent, benchmarkSize)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarkInternal(b, QuickSort, benchmarkSize)
}

func BenchmarkBubbleSort(b *testing.B) {
	benchmarkInternal(b, BubbleSort, benchmarkSize)
}

func BenchmarkHeapSort(b *testing.B) {
	benchmarkInternal(b, HeapSort, benchmarkSize)
}
