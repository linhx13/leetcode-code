package main

import "fmt"

func modifyString(s string) string {
	res := []byte{}
	var last byte
	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			res = append(res, s[i])
			last = 0
		} else {
			for c := byte('a'); c <= byte('z'); c++ {
				if (i-1 >= 0 && c == s[i-1]) || (i+1 < len(s) && c == s[i+1]) || c == last {
					continue
				}
				res = append(res, c)
				last = c
				break
			}
		}
	}
	return string(res)
}

func main() {
	// s := "?zs"
	// s := "ubv?w"
	// s := "j?qg??b"
	// s := "??yw?ipkj?"
	s := "?"
	fmt.Println(modifyString(s))
}
