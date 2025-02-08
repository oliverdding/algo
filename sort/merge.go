package sort

// MergeSort would sort the arr in place with merge sort algorithm.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) >> 1
	arr1 := MergeSort(arr[:mid])
	arr2 := MergeSort(arr[mid:])
	return merge(arr1, arr2)
}

// merge would assume the passed in array in sorted in two segment divided by mid index, and merge them to a sorted array.
func merge(arr1 []int, arr2 []int) []int {
	arr := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		arr = append(arr, arr1[i])
		i++
	}
	for j < len(arr2) {
		arr = append(arr, arr2[j])
		j++
	}
	return arr
}
