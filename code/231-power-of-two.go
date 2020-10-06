package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return ((n - 1) & n) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(218))
}
