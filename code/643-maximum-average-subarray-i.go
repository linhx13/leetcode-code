package main

func findMaxAverage(nums []int, k int) float64 {
	buf := make([]int, k)
	sum := 0
	for i := 0; i < k; i++ {
		buf[i] = nums[i]
		sum += nums[i]
	}
	res := float64(sum) / float64(k)
	idx := 0
	for i := k; i < len(nums); i++ {
		sum = sum - buf[idx] + nums[i]
		avg := float64(sum) / float64(k)
		if avg > res {
			res = avg
		}
		buf[idx] = nums[i]
		idx = (idx + 1) % k
	}
	return res
}
