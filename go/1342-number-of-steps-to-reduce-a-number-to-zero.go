package main

import "fmt"

func numberOfSteps(num int) int {
	cnt := 0
	for ; num != 0; cnt++ {
		if num&1 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
	}
	return cnt
}

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
}
