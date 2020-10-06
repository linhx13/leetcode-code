package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := []int{}
	preTime := 0
	for _, log := range logs {
		parts := strings.Split(log, ":")
		id, _ := strconv.Atoi(parts[0])
		kind := parts[1]
		time, _ := strconv.Atoi(parts[2])
		if len(stack) > 0 {
			res[stack[len(stack)-1]] += time - preTime
		}
		preTime = time
		if kind == "start" {
			stack = append(stack, id)
		} else {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[t]++
			preTime++
		}
	}
	return res
}

func main() {
	// n, logs := 2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}
	n, logs := 2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}
	fmt.Println(exclusiveTime(n, logs))
}
