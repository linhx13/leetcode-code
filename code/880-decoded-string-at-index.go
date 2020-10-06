package main

import "fmt"

func decodeAtIndex(S string, K int) string {
	strStack := []string{}
	lenStack := []int{}
	runes := []rune{}
	curLen := 0
	for _, c := range S {
		if 'a' <= c && c <= 'z' {
			runes = append(runes, c)
		} else {
			if len(runes) > 0 {
				strStack = append(strStack, string(runes))
				curLen += len(runes)
				lenStack = append(lenStack, curLen)
				runes = []rune{}
			}

			n := int(c-'1') + 1
			strStack = append(strStack, "")
			curLen = n * curLen
			lenStack = append(lenStack, curLen)
			if curLen >= K {
				break
			}
		}
	}
	if len(runes) > 0 {
		strStack = append(strStack, string(runes))
		curLen += len(runes)
		lenStack = append(lenStack, curLen)
	}
	K -= 1
	i := len(lenStack) - 1
	for i >= 0 {
		if strStack[i] == "" {
			K = K % lenStack[i-1]
			i--
		} else {
			if i == 0 {
				return string(strStack[i][K])
			} else {
				if K < lenStack[i-1] {
					i--
				} else {
					return string(strStack[i][K-lenStack[i-1]])
				}
			}
		}
	}
	return string(strStack[0][K])
}

func main() {
	// S, K := "leet2code3", 10
	// S, K := "ha22", 5
	// S, K := "a2345678999999999999999", 1
	S, K := "abc", 2
	// S, K := "a23", 6
	fmt.Println(decodeAtIndex(S, K))
}
