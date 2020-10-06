package main

import "fmt"

func helper(nums []int, cnt int, pos int, out int, res []int) []int {
	if cnt == 0 {
		res = append(res, out)
		return res
	}

	for i := pos; i < len(nums); i++ {
		res = helper(nums, cnt-1, i+1, out+nums[i], res)
	}
	return res
}

func generate(nums []int, cnt int) []int {
	var res []int
	res = helper(nums, cnt, 0, 0, res)
	return res
}

func readBinaryWatch(num int) []string {
	var res []string
	hour := []int{8, 4, 2, 1}
	minute := []int{32, 16, 8, 4, 2, 1}
	for i := 0; i <= num; i++ {
		hours := generate(hour, i)
		minutes := generate(minute, num-i)
		for _, h := range hours {
			if h > 11 {
				continue
			}
			for _, m := range minutes {
				if m > 59 {
					continue
				}
				sep := ":"
				if m < 10 {
					sep = ":0"
				}
				res = append(res, fmt.Sprintf("%d%s%d", h, sep, m))
			}
		}
	}
	return res
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
