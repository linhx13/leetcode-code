package main

import "fmt"

func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	if num == 0 {
		return res
	}
	res[1] = 1
	if num == 1 {
		return res
	}
	last := 1
	for i := 2; i <= num; i++ {
		if i == last<<1 {
			res[i] = 1
			last = i
		} else {
			res[i] = 1 + res[i^last]
		}
	}
	return res
}

func main() {
	fmt.Println(countBits(5))
}
