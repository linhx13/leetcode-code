package main

import "fmt"

func thousandSeparator(n int) string {
	res := []byte{}
	c := 0
	for n != 0 {
		if c > 0 && c%3 == 0 {
			res = append(res, '.')
		}
		res = append(res, byte('0'+n%10))
		n = n / 10
		c++
	}
	if len(res) == 0 {
		res = append(res, '0')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	// n := 123456789
	n := 0
	fmt.Println(thousandSeparator(n))
}
