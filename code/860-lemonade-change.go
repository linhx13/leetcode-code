package main

import "fmt"

func lemonadeChange(bills []int) bool {
	res_5, res_10, res_20 := 0, 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			res_5 += 1
		case 10:
			if res_5 > 0 {
				res_5 -= 1
				res_10 += 1
			} else {
				return false
			}
		case 20:
			if res_10 > 0 && res_5 > 0 {
				res_10 -= 1
				res_5 -= 1
				res_20 += 1
			} else if res_10 == 0 && res_5 > 3 {
				res_5 -= 3
				res_20 += 1
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	// bills := []int{5, 5, 5, 10, 20}
	bills := []int{5, 5, 10, 10, 20}
	fmt.Println(lemonadeChange(bills))
}
