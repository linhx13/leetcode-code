package main

func hasGroupsSizeX(deck []int) bool {
	cnts := make(map[int]int)
	for _, d := range deck {
		cnts[d]++
	}
	n := len(deck)
	for i:=2;i<=n;i++ {
		if n % i != 0 {
			continue
		}
		ok := true
		for _, v := range cnts {
			if v % i != 0 {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}
