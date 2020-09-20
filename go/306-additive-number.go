package main

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	return helper(num, 0)
}

func helper(num string, idx int) bool {
	for i := idx + 1; i < len(num); i++ {
		a, _ := strconv.Atoi(num[idx:i])
		s := strconv.Itoa(a)
		if len(s) < i-idx {
			break
		}
		for j := i + 1; j < len(num); j++ {
			a, _ := strconv.Atoi(num[idx:i])
			b, _ := strconv.Atoi(num[i:j])
			sum := strconv.Itoa(a + b)
			if len(strconv.Itoa(b)) < j-i || j+len(sum) > len(num) {
				break
			}
			if sum == num[j:j+len(sum)] {
				if j+len(sum) == len(num) || helper(num, i) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	// num := "112358"
	num := "199100199"
	fmt.Println(isAdditiveNumber(num))
}
