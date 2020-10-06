package main

import "fmt"

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n == 1 {
		return ""
	}
	runes := []rune(palindrome)
	for i := 0; i < n/2; i++ {
		if runes[i] != 'a' {
			runes[i] = 'a'
			return string(runes)
		}
	}
	for i := n - 1; i >= n/2; i-- {
		if runes[i] == 'a' {
			runes[i] = 'b'
			break
		}
	}
	return string(runes)
}

func main() {
	fmt.Println(breakPalindrome("aba"))
}
