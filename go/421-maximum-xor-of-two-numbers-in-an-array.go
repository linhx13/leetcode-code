package main

func findMaximumXOR_slow(nums []int) int {
	max := ^int(^uint(0) >> 1)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			k := nums[i] ^ nums[j]
			if k > max {
				max = k
			}
		}
	}
	return max
}

func findMaximumXOR(nums []int) int {
	res := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask |= (1 << i)
		m := make(map[int]bool)
		for _, num := range nums {
			m[num&mask] = true
		}
		t := res | (1 << i)
		for prefix, _ := range m {
			if _, ok := m[t^prefix]; ok {
				res = t
				break
			}
		}
	}
	return res
}
