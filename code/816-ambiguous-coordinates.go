package main

import "fmt"

func ambiguousCoordinates(S string) []string {
	n := len(S)
	s := S[1 : n-1]
	res := []string{}
	for i := 1; i < n-2; i++ {
		left := helper(s[0:i])
		right := helper(s[i:len(s)])
		if len(left) == 0 || len(right) == 0 {
			continue
		}
		for _, l := range left {
			for _, r := range right {
				res = append(res, fmt.Sprintf("(%s, %s)", l, r))
			}
		}
	}
	return res
}

func helper(s string) []string {
	n := len(s)
	if n == 1 {
		return []string{s}
	}
	if s[0] == '0' && s[n-1] == '0' {
		return nil
	}
	if s[0] == '0' {
		return []string{fmt.Sprintf("0.%s", s[1:n])}
	}
	if s[n-1] == '0' {
		return []string{s}
	}
	res := []string{s}
	for i := 1; i < n; i++ {
		res = append(res, fmt.Sprintf("%s.%s", s[0:i], s[i:n]))
	}
	return res
}

func main() {
	S := "(123)"
	// S := "(00011)"
	// S := "(100)"
	fmt.Println(ambiguousCoordinates(S))
}
