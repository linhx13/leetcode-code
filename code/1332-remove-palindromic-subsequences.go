package main

import "fmt"

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	isPal := true
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			isPal = false
			break
		}
	}
	if isPal {
		return 1
	}
	return 2
}

func main() {
	s := ""
	fmt.Println(removePalindromeSub(s))
}
