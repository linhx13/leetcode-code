package main

func totalHammingDistance(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, num := range nums {
			if num&(1<<i) != 0 {
				cnt++
			}
		}
		res += cnt * (n - cnt)
	}
	return res
}
