package main

func sortString(s string) string {
	counts := make([]int, 26)
	total := 0
	for _, r := range s {
		total++
		counts[r-'a']++
	}
	res := []rune{}
	for total > 0 {
		for i := 0; i < 26; i++ {
			if counts[i] > 0 {
				res = append(res, rune('a'+i))
				counts[i]--
				total--
			}
		}
		for i := 26 - 1; i >= 0; i-- {
			if counts[i] > 0 {
				res = append(res, rune('a'+i))
				counts[i]--
				total--
			}
		}
	}
	return string(res)
}
