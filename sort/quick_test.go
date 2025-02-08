// Package sort contains algorithms that related to sorting.
package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert.Equal(t, []int{}, QuickSort([]int{}), "empty slice")
	assert.Equal(t, []int{1}, QuickSort([]int{1}), "single element slice")
	assert.Equal(t, []int{1, 2}, QuickSort([]int{2, 1}), "two elements slice")
	assert.Equal(t, []int{1, 2, 3, 4, 5}, QuickSort([]int{5, 3, 1, 2, 4}), "simple slice")
	assert.Equal(t, []int{1, 2}, QuickSort([]int{1, 2}), "sorted slice")
	assert.Equal(t, []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, QuickSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10}), "reversed slice")
}

func BenchmarkQuickSort1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(1)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(10)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(100)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(1000)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(10000)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(100000)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(1000000)
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort10000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(10000000)
		b.StartTimer()
		QuickSort(arr)
	}
}
