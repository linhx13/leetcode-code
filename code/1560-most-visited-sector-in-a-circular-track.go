package main

import (
	"fmt"
	"sort"
)

func mostVisited(n int, rounds []int) []int {
	cnts := make(map[int]int)
	s := rounds[0]
	for i := 1; i < len(rounds); i++ {
		for {
			cnts[s]++
			s++
			if s > n {
				s = 1
			}
			if s == rounds[i] {
				break
			}
		}
	}
	cnts[s]++
	max := ^int(^uint(0) >> 1)
	for _, v := range cnts {
		if v > max {
			max = v
		}
	}
	res := []int{}
	for k, v := range cnts {
		if v == max {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}

func main() {
	// n := 4
	// rounds := []int{1, 3, 1, 2}
	// n := 2
	// rounds := []int{2, 1, 2, 1, 2, 1, 2, 1, 2}
	n := 7
	rounds := []int{1, 3, 5, 7}
	fmt.Println(mostVisited(n, rounds))
}
