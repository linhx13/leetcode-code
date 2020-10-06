package main

func prefixesDivBy5(A []int) []bool {
	num := 0
	res := make([]bool, len(A))
	for i, n := range A {
		num = ((num << 1) + n) % 5
		if num == 0 {
			res[i] = true
		}
	}
	return res
}
