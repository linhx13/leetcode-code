package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	res := 24*60 + 1
	for i := 1; i < len(timePoints); i++ {
		cur := diff(timePoints[i-1], timePoints[i])
		if cur < res {
			res = cur
		}
	}
	cur := diff(timePoints[len(timePoints)-1], timePoints[0]) + 24*60
	if cur < res {
		res = cur
	}
	return res
}

func diff(s, t string) int {
	s_arr := strings.Split(s, ":")
	t_arr := strings.Split(t, ":")
	s1, _ := strconv.Atoi(s_arr[0])
	s2, _ := strconv.Atoi(s_arr[1])
	t1, _ := strconv.Atoi(t_arr[0])
	t2, _ := strconv.Atoi(t_arr[1])
	return (t1-s1)*60 + (t2 - s2)
}

func main() {
	// timePoints := []string{"23:59", "22:33", "00:00"}
	timePoints := []string{"23:59", "22:30", "12:00"}
	fmt.Println(findMinDifference(timePoints))
}
