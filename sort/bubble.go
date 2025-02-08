package sort

// BubbleSort would sort the arr in place with bubble sort algorithm.
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
