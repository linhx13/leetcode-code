package main

func numberOfLines(widths []int, S string) []int {
	lines, units := 1, 0
	for _, c := range S {
		if units+widths[c-'a'] > 100 {
			lines++
			units = 0
		}
		units += widths[c-'a']
	}
	return []int{lines, units}
}
