// Package search contains algorithms that related to searching.
package search

// BinarySearchLast would return the last position of target element in the array which could contain duplicate elements.
// - If the element is found, the first return int is its position, and the second return int is -1.
// - If the element is not found, the first return int is -1, and the second return int is its position should be inserted.
func BinarySearchLast(arr []int, target int) (int, int) {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + ((hi - lo) >> 2)
		if arr[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if lo > 0 && arr[lo-1] == target {
		return lo - 1, -1
	}
	return -1, lo
}
