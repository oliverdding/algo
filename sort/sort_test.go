package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func generateRandomArray(arrayLen int) []int {
	arr := make([]int, 0, arrayLen)
	for i := 0; i < arrayLen; i++ {
		arr = append(arr, rand.Int())
	}
	return arr
}

func BenchmarkStableSort100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(100000)
		b.StartTimer()
		sort.SliceStable(arr, func(i, j int) bool {
			return i > j
		})
	}
}

func BenchmarkStableSort1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(1000000)
		b.StartTimer()
		sort.SliceStable(arr, func(i, j int) bool {
			return i > j
		})
	}
}

func BenchmarkStableSort10000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		arr := generateRandomArray(10000000)
		b.StartTimer()
		sort.SliceStable(arr, func(i, j int) bool {
			return i > j
		})
	}
}
