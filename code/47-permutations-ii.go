package main

import (
	"fmt"
	"strconv"
	"strings"
)

func permuteUnique(nums []int) [][]int {
	cnts := make(map[int]int)
	for _, n := range nums {
		cnts[n]++
	}
	res := make(map[string][]int)
	helper(cnts, &[]int{}, res)
	vals := [][]int{}
	for _, v := range res {
		vals = append(vals, v)
	}
	return vals
}

func helper(cnts map[int]int, buf *[]int, res map[string][]int) {
	flag := false
	for i, n := range cnts {
		if n == 0 {
			continue
		}
		flag = true
		cnts[i]--
		*buf = append(*buf, i)
		helper(cnts, buf, res)
		*buf = (*buf)[0 : len(*buf)-1]
		cnts[i]++
	}
	if !flag {
		ints := make([]int, len(*buf))
		copy(ints, *buf)
		keys := []string{}
		for _, i := range ints {
			keys = append(keys, strconv.Itoa(i))
		}
		res[strings.Join(keys, "#")] = ints
	}
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
