package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] >= target:
			high = mid
		}
	}

	return -1
}

func main() {
	b := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(b)
	fmt.Println(search(b, 12))
}
