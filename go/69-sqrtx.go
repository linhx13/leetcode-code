package main

import "fmt"

func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := low + (high-low)>>1
		n := mid * mid
		if n <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if n >= x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(mySqrt(9))
}
