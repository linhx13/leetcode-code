package main

import "fmt"

func findComplement(num int) int {
	start := false
	for i := 32; i >= 0; i-- {
		if num&(1<<i) != 0 {
			start = true
		}
		if start {
			num ^= (1 << i)
		}
	}
	return num
}

func main() {
	fmt.Println(findComplement(1))
}
