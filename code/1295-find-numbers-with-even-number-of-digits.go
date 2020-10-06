package main

func findNumbers(nums []int) int {
	res := 0
	for _, n := range nums {
		c := 0
		for n > 0 {
			c++
			n = n / 10
		}
		if c%2 == 0 {
			res++
		}
	}
	return res
}
