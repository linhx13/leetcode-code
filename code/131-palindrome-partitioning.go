package main

import "fmt"

func partition(s string) [][]string {
	res := [][]string{}
	helper(s, &[]string{}, &res)
	return res
}

func helper(s string, buf *[]string, res *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(*buf))
		copy(tmp, *buf)
		*res = append(*res, tmp)
		return
	}

	for i := 1; i <= len(s); i++ {
		flag := true
		for j := 0; j < i; j++ {
			if s[j] != s[i-j-1] {
				flag = false
				break
			}
		}
		if flag {
			*buf = append(*buf, s[0:i])
			helper(s[i:len(s)], buf, res)
			*buf = (*buf)[0 : len(*buf)-1]
		}
	}
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
