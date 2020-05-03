package main

import "fmt"

func commonChars(A []string) []string {
	gm := make(map[rune]int)
	for i, s := range A {
		cm := make(map[rune]int)
		for _, c := range s {
			cm[c] = cm[c] + 1
		}
		if i == 0 {
			gm = cm
		} else {
			nm := make(map[rune]int)
			for k, cv := range cm {
				gv := gm[k]
				if cv < gv {
					gv = cv
				}
				if gv > 0 {
					nm[k] = gv
				}
			}
			gm = nm
		}
	}
	res := []string{}
	for k, v := range gm {
		for i := 0; i < v; i++ {
			res = append(res, string(k))
		}
	}
	return res
}

func main() {
	// A := []string{"bella", "label", "roller"}
	// A := []string{"cool", "lock", "cook"}
	A := []string{"aaa"}
	fmt.Println(commonChars(A))
}
