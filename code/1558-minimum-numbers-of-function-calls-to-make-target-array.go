package main

import "fmt"

func minOperations(nums []int) int {
	res := 0
	max := 0
	for _, n := range nums {
		c := 0
		for n != 0 {
			if n&1 == 1 {
				res++
			}
			n >>= 1
			c++
		}
		if c-1 > max {
			max = c - 1
		}
	}
	return res + max
}

func main() {
	// nums := []int{1, 5}
	nums := []int{2, 4, 8, 16}
	fmt.Println(minOperations(nums))
}
