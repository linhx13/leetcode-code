package main

import "fmt"

func balancedStringSplit(s string) int {
	cnt, sum := 0, 0
	for _, c := range s {
		if c == 'L' {
			sum++
		} else {
			sum--
		}
		if sum == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	// s := "RLRRLLRLRL"
	// s := "RLLLLRRRLR"
	// s := "LLLLRRRR"
	s := "RLRRRLLRLL"
	fmt.Println(balancedStringSplit(s))
}
