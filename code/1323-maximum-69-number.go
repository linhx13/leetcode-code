package main

import "fmt"

func maximum69Number(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	found := false
	res := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 6 && !found {
			digits[i] = 9
			found = true
		}
		res = res*10 + digits[i]
	}
	return res
}

func main() {
	fmt.Println(maximum69Number(696))
}
