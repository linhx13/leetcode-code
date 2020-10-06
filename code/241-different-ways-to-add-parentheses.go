package main

import (
	"fmt"
	"sort"
)

func diffWaysToCompute(input string) []int {
	cache := make(map[string][]int)
	res := helper(input, 0, len(input), cache)
	sort.Ints(res)
	return res
}

func helper(input string, start, end int, cache map[string][]int) []int {
	key := fmt.Sprintf("%d-%d", start, end)
	if val, ok := cache[key]; ok {
		return val
	}
	res := []int{}
	for i := start; i < end; i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := helper(input, start, i, cache)
			right := helper(input, i+1, end, cache)
			for _, x := range left {
				for _, y := range right {
					if input[i] == '+' {
						res = append(res, x+y)
					} else if input[i] == '-' {
						res = append(res, x-y)
					} else if input[i] == '*' {
						res = append(res, x*y)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		num := 0
		for i := start; i < end; i++ {
			if '0' <= input[i] && input[i] <= '9' {
				num = num*10 + int(input[i]-'0')
			}
		}
		res = append(res, num)
	}
	cache[key] = res
	return res
}

func main() {
	// input := "2-1-1"
	// input := "2*3-4*5"
	input := "11"
	fmt.Println(diffWaysToCompute(input))
}
