// Package search contains algorithms that related to searching.
package search

// BinarySearch would return the position of target element in the array which contains no duplicate elements.
// - If the element is found, the first return int is its position, and the second return int is -1.
// - If the element is not found, the first return int is -1, and the second return int is its position should be inserted.
func BinarySearch(arr []int, target int) (int, int) {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + ((hi - lo) >> 2)
		if arr[mid] > target {
			hi = mid
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			return mid, -1
		}
	}
	return -1, lo
}
