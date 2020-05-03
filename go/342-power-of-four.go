package main

import "fmt"

// func isPowerOfFour(num int) bool {
// 	return num > 0 && (num&(num-1)) == 0 && (num&0x55555555) == num
// }

func isPowerOfFour(num int) bool {
	if num == 1 {
		return true
	}
	if num%2 != 0 || num%4 != 0 {
		return false
	}
	for num != 0 && (num&0x3) == 0 {
		num >>= 2
	}
	return num == 1
}

func main() {
	fmt.Println(isPowerOfFour(16))
}
