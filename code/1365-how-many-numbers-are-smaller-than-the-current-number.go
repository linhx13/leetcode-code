package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	counter := make(map[int]int)
	for _, n := range nums {
		counter[n] += 1
	}
	for i := 1; i <= 100; i++ {
		counter[i] = counter[i] + counter[i-1]
	}
	res := []int{}
	for _, n := range nums {
		res = append(res, counter[n-1])
	}
	return res
}

func main() {
	// nums := []int{6, 5, 4, 8}
	// nums := []int{7, 7, 7, 7}
	nums := []int{1}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
