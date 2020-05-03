package main

import "fmt"

func replaceElements(arr []int) []int {
	old_max, new_max := -1, -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > new_max {
			new_max = arr[i]
		}
		arr[i] = old_max
		old_max = new_max
	}
	return arr
}

func main() {
	// arr := []int{17, 18, 5, 4, 6, 1}
	// arr := []int{1, 1, 1, 1, 1}
	arr := []int{1}
	fmt.Println(replaceElements(arr))
}
