package main

import "fmt"

func calculate(s string) int {
	nums := []int{}
	ops := []byte{}
	i := 0
	for i < len(s) {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			ops = append(ops, s[i])
		} else if '0' <= s[i] && s[i] <= '9' {
			cur := 0
			for ;i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
				cur = cur * 10 + int(s[i] - '0')
			}
			i--
			if len(ops) > 0 && ops[len(ops)-1] == '*' {
				nums[len(nums)-1] = nums[len(nums)-1] * cur
				ops = ops[0:len(ops)-1]
			} else if len(ops) > 0 && ops[len(ops)-1] == '/'{
				nums[len(nums)-1] = nums[len(nums)-1] / cur
				ops = ops[0:len(ops)-1]
			} else {
				nums = append(nums, cur)
			}
		}
		i++
	}
	if len(ops) == 0 && len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 0; i < len(ops); i++ {
		if ops[i] == '+' {
			res += nums[i+1]
		} else {
			res -= nums[i+1]
		}
	}
	return res
}

func main() {
	// s := "3+2*2"
	// s := " 3  /2"
	s := " 3+15 / 2 -1"
	// s := "1-1+1"
	fmt.Println(calculate(s))
}
