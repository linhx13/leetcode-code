package main

func countCharacters(words []string, chars string) int {
	chars_cnt := make(map[rune]int)
	for _, c := range chars {
		chars_cnt[c] = chars_cnt[c] + 1
	}
	res := 0
	for _, word := range words {
		word_cnt := make(map[rune]int)
		for _, c := range word {
			word_cnt[c] = word_cnt[c] + 1
		}
		isValid := true
		for c, v := range word_cnt {
			char_v := chars_cnt[c]
			if v > char_v {
				isValid = false
				break
			}
		}
		if isValid {
			res += len(word)
		}
	}
	return res
}
