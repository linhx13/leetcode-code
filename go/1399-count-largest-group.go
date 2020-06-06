package main

func countLargestGroup(n int) int {
	sum_cnts := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum := 0
		for x := i; x > 0; {
			sum += x % 10
			x = x / 10
		}
		sum_cnts[sum] = sum_cnts[sum] + 1
	}
	max_cnt := 0
	for _, v := range sum_cnts {
		if v > max_cnt {
			max_cnt = v
		}
	}
	res := 0
	for _, v := range sum_cnts {
		if v == max_cnt {
			res++
		}
	}
	return res
}
