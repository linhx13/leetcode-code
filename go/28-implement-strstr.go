package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	lh, ln := len(haystack), len(needle)
	for i := 0; i < lh; i++ {
		if haystack[i] != needle[0] {
			continue
		}
		if len(haystack)-i < len(needle) {
			return -1
		}
		j := 0
		for ; j < ln && i+j < lh && haystack[i+j] == needle[j]; j++ {
		}
		if j == ln {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("allo", ""))
}
