package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maximumProduct(nums []int) int {
	max1, max2, max3 := INT_MIN, INT_MIN, INT_MIN
	min1, min2 := INT_MAX, INT_MAX

	for _, n := range nums {
		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 {
			max3 = max2
			max2 = n
		} else if n > max3 {
			max3 = n
		}
		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		}
	}
	x := max1 * max2 * max3
	y := max1 * min1 * min2
	if y > x {
		x = y
	}
	return x
}

func main() {
	nums := []int{-4, -3, -2, -1, 60}
	fmt.Println(maximumProduct(nums))
}
