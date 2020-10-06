package main

func maxProduct(nums []int) int {
	m1, m2 := 0, 0
	for _, n := range nums {
		if n > m1 {
			m2 = m1
			m1 = n
		} else if n > m2 {
			m2 = n
		}
	}
	return (m1 - 1) * (m2 - 1)
}
