package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i]&1 == 0 {
			continue
		}
		if i-1 >= 0 && arr[i-1]&1 == 1 && i+1 < len(arr) && arr[i+1]&1 == 1 {
			return true
		}
	}
	return false
}

func main() {
	// arr := []int{2, 6, 4, 1}
	// arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	arr := []int{1, 1, 1}
	fmt.Println(threeConsecutiveOdds(arr))
}
