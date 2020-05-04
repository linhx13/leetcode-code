package main

import "fmt"

func freqAlphabets(s string) string {
	res := []rune{}
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		x := 'a'
		if runes[i] == '#' {
			x += (runes[i-2]-'0')*10 + (runes[i-1] - '0') - 1
			i -= 2
		} else {
			x += runes[i] - '0' - 1
		}
		res = append(res, x)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}

func main() {
	// s := "10#11#12"
	// s := "1326#"
	// s := "25#"
	// s := "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
	s := ""
	fmt.Println(freqAlphabets(s))
}
