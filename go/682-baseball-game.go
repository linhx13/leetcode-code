package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	res := 0
	scores := []int{}
	for _, op := range ops {
		switch op {
		case "+":
			s := scores[len(scores)-1] + scores[len(scores)-2]
			scores = append(scores, s)
			res += s
		case "D":
			s := 2 * scores[len(scores)-1]
			scores = append(scores, s)
			res += s
		case "C":
			s := scores[len(scores)-1]
			scores = scores[:len(scores)-1]
			res -= s
		default:
			s, _ := strconv.Atoi(op)
			scores = append(scores, s)
			res += s
		}
	}
	return res
}

func main() {
	// ops := []string{"5", "2", "C", "D", "+"}
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(ops))
}
