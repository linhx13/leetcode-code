package main

import "fmt"

func predictPartyVictory(senate string) string {
	cnt := 0
	seenR, seenD := true, true
	ss := []rune(senate)
	for seenR && seenD {
		seenR, seenD = false, false
		for i := 0; i < len(ss); i++ {
			if ss[i] == 'R' {
				if cnt < 0 {
					ss[i] = 'O'
				}
				cnt++
				seenR = true
			} else if ss[i] == 'D' {
				if cnt > 0 {
					ss[i] = 'O'
				}
				cnt--
				seenD = true
			}
		}
	}
	if seenR {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func main() {
	// senate := "RD"
	senate := "RDD"
	fmt.Println(predictPartyVictory(senate))
}
