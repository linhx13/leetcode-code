package main

import "fmt"

func longestDiverseString(a int, b int, c int) string {
	res := []rune{}
	total := a + b + c
	curA, curB, curC := 0, 0, 0
	for total > 0 {
		if (a >= b && a >= c && curA != 2) || (a > 0 && (curB == 2 || curC == 2)) {
			res = append(res, 'a')
			a--
			curA++
			curB, curC = 0, 0
		} else if (b >= a && b >= c && curB != 2) || (b > 0 && (curA == 2 || curC == 2)) {
			res = append(res, 'b')
			b--
			curB++
			curA, curC = 0, 0
		} else if (c >= a && c >= b && curC != 2) || (c > 0 && (curA == 2 || curB == 2)) {
			res = append(res, 'c')
			c--
			curC++
			curA, curB = 0, 0
		}
		total--
	}
	return string(res)
}

func main() {
	a, b, c := 1, 1, 7
	fmt.Println(longestDiverseString(a, b, c))
}
