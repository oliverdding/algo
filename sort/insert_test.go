package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	assert.Equal(t, []int{}, InsertSort([]int{}), "empty slice")
	assert.Equal(t, []int{1}, InsertSort([]int{1}), "single element slice")
	assert.Equal(t, []int{1, 2}, InsertSort([]int{2, 1}), "two elements slice")
	assert.Equal(t, []int{1, 2, 3, 4, 5}, InsertSort([]int{5, 3, 1, 2, 4}), "simple slice")
	assert.Equal(t, []int{1, 2}, InsertSort([]int{1, 2}), "sorted slice")
	assert.Equal(t, []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, InsertSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10}), "reversed slice")
}

func BenchmarkInsertSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(100000)
		b.StartTimer()
		InsertSort(arr)
	}
}
