package main

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	res := []string{}
	res = append(res, strconv.Itoa(nums[0]))
	if len(nums) == 2 {
		res = append(res, "/", strconv.Itoa(nums[1]))
	} else if len(nums) > 2 {
		res = append(res, "/(")
		res = append(res, strconv.Itoa(nums[1]))
		for i := 2; i < len(nums); i++ {
			res = append(res, "/", strconv.Itoa(nums[i]))
		}
		res = append(res, ")")
	}

	return strings.Join(res, "")
}

func main() {
	nums := []int{1000, 100, 10, 2}
	// nums := []int{1, 10}
	fmt.Println(optimalDivision(nums))
}
