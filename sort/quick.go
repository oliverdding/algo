// Package sort contains algorithms that related to sorting.
package sort

// QuickSort would sort the arr in place with quick sort algoritm.
func QuickSort(arr []int) []int {
	if len(arr) > 1 {
		pivot := partition(arr)
		QuickSort(arr[0:pivot])
		QuickSort(arr[pivot+1:])
	}
	return arr
}

func partition(arr []int) uint {
	lo, base := 0, len(arr)-1
	for i := 0; i < base; i++ {
		if arr[i] < arr[base] {
			arr[lo], arr[i] = arr[i], arr[lo]
			lo++
		}
	}
	arr[lo], arr[base] = arr[base], arr[lo]
	return uint(lo)
}
