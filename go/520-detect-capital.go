package main

func detectCapitalUse(word string) bool {
	cnt := 0
	for i, r := range word {
		if 'A' <= r && r <= 'Z' {
			if i > 0 && cnt == 0 {
				return false
			}
			cnt++
		}
	}
	return cnt == 0 || cnt == len(word) || cnt == 1
}
