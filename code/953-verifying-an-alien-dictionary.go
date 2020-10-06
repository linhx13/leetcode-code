package main

func isAlienSorted(words []string, order string) bool {
	if len(words) == 1 {
		return true
	}
	rank := make(map[rune]int)
	for i, r := range order {
		rank[r] = i
	}

	for i := 1; i < len(words); i++ {
		if !check(words[i-1], words[i], rank) {
			return false
		}
	}
	return true
}

func check(a string, b string, rank map[rune]int) bool {
	ra, rb := []rune(a), []rune(b)
	for i := 0; i < len(ra) && i < len(rb); i++ {
		if rank[ra[i]] < rank[rb[i]] {
			return true
		} else if rank[ra[i]] > rank[rb[i]] {
			return false
		}
	}
	return len(ra) <= len(rb)
}
