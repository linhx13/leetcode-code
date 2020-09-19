package main

import "fmt"

func singleNumber(nums []int) int {
	one, two, three := 0, 0, 0
	for _, n := range nums {
		two |= one & n
		one ^= n
		three = one & two
		one &= ^three
		two &= ^three
	}
	return one
}

func main() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))
}
