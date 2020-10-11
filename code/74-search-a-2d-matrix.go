package main

func searchMatrix(matrix [][]int, target int) bool {
	arr := []int{}
	for _, row := range matrix {
		arr = append(arr, row...)
	}
	low, high := 0, len(arr)
	for low < high {
		mid := low + (high - low) >> 1
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return false
}
