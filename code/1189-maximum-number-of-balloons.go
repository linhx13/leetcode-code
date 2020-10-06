package main

func maxNumberOfBalloons(text string) int {
	cnts := make(map[rune]int)
	for _, r := range text {
		cnts[r]++
	}
	m := map[rune]int{
		'b': 1,
		'a': 1,
		'l': 2,
		'o': 2,
		'n': 1,
	}
	min := 99999
	for k, v := range m {
		c := cnts[k] / v
		if c < min {
			min = c
		}
	}
	return min
}
