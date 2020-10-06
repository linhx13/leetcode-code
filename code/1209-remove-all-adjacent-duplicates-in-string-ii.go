package main

import "fmt"

func removeDuplicates(s string, k int) string {
	stack := []rune{}
	for _, c := range s {
		stack = append(stack, c)
		for {
			n := len(stack)
			if n < k {
				break
			}
			valid := true
			for i := 1; i <= k; i++ {
				if stack[n-i] != stack[n-1] {
					valid = false
					break
				}
			}
			if valid {
				stack = stack[:n-k]
			} else {
				break
			}
		}
	}
	return string(stack)
}

func main() {
	// s, k := "abcd", 2
	// s, k := "deeedbbcccbdaa", 3
	// s, k := "pbbcggttciiippooaais", 2
	s, k := "aa", 2
	fmt.Println(removeDuplicates(s, k))
}
