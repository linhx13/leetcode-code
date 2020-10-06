package main

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	re := regexp.MustCompile(`[!?',;.]+`)
	text := re.ReplaceAllLiteralString(paragraph, " ")
	words := strings.Fields(strings.ToLower(text))
	bannedMap := make(map[string]bool)
	for _, s := range banned {
		bannedMap[s] = true
	}
	freqs := make(map[string]int)
	for _, word := range words {
		if !bannedMap[word] {
			freqs[word]++
		}
	}
	maxCnt, maxWord := 0, ""
	for word, cnt := range freqs {
		if cnt > maxCnt {
			maxWord = word
			maxCnt = cnt
		}
	}
	return maxWord
}

func main() {

}
