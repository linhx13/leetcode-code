package main

func sumZero(n int) []int {
	res := []int{}
	for i := 1; i*2 <= n; i++ {
		res = append(res, i, -i)
	}
	if n%2 == 1 {
		res = append(res, 0)
	}
	return res
}
