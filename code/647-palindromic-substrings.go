package main

import "fmt"

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		helper(s, i, i, &res)
		helper(s, i, i+1, &res)
	}
	return res
}

func helper(s string, i, j int, res *int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		*res++
	}
}

func main() {
	// s := "abc"
	s := "aaa"
	fmt.Println(countSubstrings(s))
}
