package sort

import "math/rand"

func generateRandomArray(arrayLen int) []int {
	arr := make([]int, 0, arrayLen)
	for i := 0; i < arrayLen; i++ {
		arr = append(arr, rand.Int())
	}
	return arr
}
