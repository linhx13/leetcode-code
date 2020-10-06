package main

import "fmt"

func getWinner(arr []int, k int) int {
	max := arr[0]
	if arr[1] > max {
		max = arr[1]
	}
	cnt := 1
	for i := 2; i < len(arr); i++ {
		if cnt >= k {
			return max
		}
		if arr[i] > max {
			max, cnt = arr[i], 1
		} else {
			cnt++
		}
	}
	return max
}

func main() {
	arr := []int{2, 1, 3, 5, 4, 6, 7}
	k := 2
	// arr := []int{1, 9, 8, 2, 3, 7, 6, 4, 5}
	// k := 7
	// arr := []int{3, 2, 1}
	// k := 10
	// arr := []int{1, 25, 35, 42, 68, 70}
	// k := 1
	fmt.Println(getWinner(arr, k))
}
