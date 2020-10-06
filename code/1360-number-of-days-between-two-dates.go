package main

import (
	"math"
	"strconv"
)

func daysBetweenDates(date1 string, date2 string) int {
	return int(math.Abs(float64(daysFromEpoch(date1) - daysFromEpoch(date2))))
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func daysFromEpoch(date string) int {
	m := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[5:7])
	days, _ := strconv.Atoi(date[8:10])
	for i := 1970; i < year; i++ {
		days += 365
		if isLeap(i) {
			days++
		}
	}
	for i := 1; i < month; i++ {
		days += m[i-1]
	}
	if month > 2 && isLeap(year) {
		days++
	}
	return days
}
