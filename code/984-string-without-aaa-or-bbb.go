package main

import "fmt"

func strWithout3a3b(A int, B int) string {
	res := []byte{}
	cnt1, cnt2 := A, B
	ch1, ch2 := byte('a'), byte('b')
	if cnt1 < cnt2 {
		cnt1, cnt2 = cnt2, cnt1
		ch1, ch2 = ch2, ch1
	}
	for cnt1-cnt2 > 0 {
		for i := 0; i < 2 && cnt1 > 0; i++ {
			res = append(res, ch1)
			cnt1--
		}
		if cnt2 > 0 {
			res = append(res, ch2)
			cnt2--
		}
	}
	for cnt1 > 0 || cnt2 > 0 {
		if cnt1 > 0 {
			res = append(res, ch1)
			cnt1--
		}
		if cnt2 > 0 {
			res = append(res, ch2)
			cnt2--
		}
	}
	return string(res)
}

func main() {
	// A, B := 1, 2
	// A, B := 4, 1
	A, B := 5, 3
	fmt.Println(strWithout3a3b(A, B))
}
