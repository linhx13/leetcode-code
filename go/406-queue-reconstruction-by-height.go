package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})
	res := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		j, c := 0, 0
		for ; j < i && c < people[i][1]; j++ {
			if res[j][0] >= people[i][0] {
				c++
			}
		}
		for k := i - 1; k >= j; k-- {
			res[k+1] = res[k]
		}
		res[j] = people[i]
	}
	return res
}

func main() {
	people := [][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0}, []int{6, 1}, []int{5, 2}}
	fmt.Println(reconstructQueue(people))
}
