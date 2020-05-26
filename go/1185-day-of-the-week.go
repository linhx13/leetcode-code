package main

func leapYear(y int) bool {
	return (y%4 == 0 && y%100 != 0) || (y%400 == 0)
}

func dayOfTheWeek(day int, month int, year int) string {
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if leapYear(year) {
		days[1]++
	}
	sum := 0
	for i := 1970; i < year; i++ {
		sum += 365
		if leapYear(i) {
			sum++
		}
	}
	for i := 1; i < month; i++ {
		sum += days[i-1]
	}
	sum += day
	return names[(sum+3)%7]
}
