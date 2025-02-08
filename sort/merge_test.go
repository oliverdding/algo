package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	assert.Equal(t, []int{}, MergeSort([]int{}), "empty slice")
	assert.Equal(t, []int{1}, MergeSort([]int{1}), "single element slice")
	assert.Equal(t, []int{1, 2}, MergeSort([]int{2, 1}), "two elements slice")
	assert.Equal(t, []int{1, 2, 3, 4, 5}, MergeSort([]int{5, 3, 1, 2, 4}), "simple slice")
	assert.Equal(t, []int{1, 2}, MergeSort([]int{1, 2}), "sorted slice")
	assert.Equal(t, []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, MergeSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10}), "reversed slice")
}

func BenchmarkMergeSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(100000)
		b.StartTimer()
		MergeSort(arr)
	}
}

func BenchmarkMergeSort1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(1000000)
		b.StartTimer()
		MergeSort(arr)
	}
}

func BenchmarkMergeSort10000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(10000000)
		b.StartTimer()
		MergeSort(arr)
	}
}
