package main

func isSubsequence(s string, t string) bool {
	n := len(t)
	i := 0
	runes := []rune(t)
	for _, c := range s {
		for i < n && runes[i] != c {
			i++
		}
		if i == n {
			return false
		}
		i++
	}
	return true
}
