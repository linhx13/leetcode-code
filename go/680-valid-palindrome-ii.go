package main

import "fmt"

func isValid(s string, left int, right int) bool {
	for ; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return isValid(s, left, right-1) || isValid(s, left+1, right)
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abaa"))
}
