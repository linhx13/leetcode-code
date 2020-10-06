package main

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	arr := strings.Split(date, "-")
	year, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	day, _ := strconv.Atoi(arr[2])
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	res := 0
	for i := 0; i < month-1; i++ {
		res += days[i]
	}
	res += day
	if isLeap(year) && month > 2 {
		res += 1
	}
	return res
}

func isLeap(y int) bool {
	return (y%4 == 0 && y%100 != 0) || y%400 == 0
}
