package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reformatDate(date string) string {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	monthMap := make(map[string]string)
	for i, m := range months {
		monthMap[m] = fmt.Sprintf("%02d", i+1)
	}
	fields := strings.Fields(date)
	dayN, _ := strconv.Atoi(fields[0][:len(fields[0])-2])
	day := fmt.Sprintf("%02d", dayN)
	month := monthMap[fields[1]]
	year := fields[2]
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
