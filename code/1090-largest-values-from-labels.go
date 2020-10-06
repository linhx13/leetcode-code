package main

import (
	"fmt"
	"sort"
)

type Item struct {
	value int
	label int
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	items := []Item{}
	for i := 0; i < len(values); i++ {
		item := Item{value: values[i], label: labels[i]}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	res := 0
	num_used := 0
	label_used := make(map[int]int)
	for i := 0; i < len(items) && num_used < num_wanted; i++ {
		lc := label_used[items[i].label]
		if lc >= use_limit {
			continue
		}
		res += items[i].value
		num_used++
		label_used[items[i].label]++
	}
	return res
}

func main() {
	// values := []int{5, 4, 3, 2, 1}
	// labels := []int{1, 1, 2, 2, 3}
	// num_wanted := 3
	// use_limit := 1
	values := []int{9, 8, 8, 7, 6}
	labels := []int{0, 0, 0, 1, 1}
	num_wanted := 3
	use_limit := 1
	fmt.Println(largestValsFromLabels(values, labels, num_wanted, use_limit))
}
