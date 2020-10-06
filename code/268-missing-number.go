package main

import "fmt"

func missingNumber(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	all := (1 + len(nums)) * len(nums) / 2
	return all - sum
}

func main() {
	// nums := []int{3, 0, 1}
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
