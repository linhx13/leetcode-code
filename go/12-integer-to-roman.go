package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	m := make(map[int]string)
	m[1] = "I"
	m[5] = "V"
	m[10] = "X"
	m[50] = "L"
	m[100] = "C"
	m[500] = "D"
	m[1000] = "M"
	m[4] = "IV"
	m[9] = "IX"
	m[40] = "XL"
	m[90] = "XC"
	m[400] = "CD"
	m[900] = "CM"

	res := []string{}
	base := 1
	for num > 0 {
		cur := num % 10
		if cur == 4 || cur == 9 || cur == 5 {
			res = append(res, m[cur*base])
		} else if cur < 4 {
			for i := 1; i <= cur; i++ {
				res = append(res, m[base])
			}
		} else {
			for i := 1; i <= cur-5; i++ {
				res = append(res, m[base])
			}
			res = append(res, m[5*base])
		}
		num = num / 10
		base *= 10
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, "")
}

func main() {
	num := 1994
	fmt.Println(intToRoman(num))
}
