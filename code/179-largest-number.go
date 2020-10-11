package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	arr := []string{}
	for _, n := range nums {
		arr = append(arr, strconv.Itoa(n))
	}
	sort.Slice(arr, func(i, j int) bool {
		return cmp(arr[i], arr[j])
	})
	res := strings.Join(arr, "")
	if res[0] == '0' {
		i := 1
		for ; i < len(res) && res[i] == '0'; i++ {
		}
		res = res[i-1 : len(res)]
	}
	return res
}

func cmp(x, y string) bool {
	a := x + y
	b := y + x
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}

func main() {
	// nums := []int{3, 30, 34}
	// nums := []int{10, 2}
	// nums := []int{3, 30, 34, 5, 9}
	// nums := []int{10}
	// nums := []int{34323, 3432}
	// nums := []int{353, 35}
	nums := []int{0, 0}
	fmt.Println(largestNumber(nums))
}
