package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, c := range candies {
		if c > max {
			max = c
		}
	}
	res := []bool{}
	for _, c := range candies {
		r := false
		if c+extraCandies >= max {
			r = true
		}
		res = append(res, r)
	}
	return res
}

func main() {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
