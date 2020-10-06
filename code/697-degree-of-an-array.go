package main

import "fmt"

type Number struct {
	Freq  int
	Start int
	End   int
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]Number)
	da := 0
	for i, n := range nums {
		num, ok := m[n]
		if ok {
			num.Freq += 1
			num.End = i
			m[n] = num
		} else {
			num = Number{}
			num.Freq = 1
			num.Start = i
			num.End = i
			m[n] = num
		}
		if m[n].Freq > da {
			da = m[n].Freq
		}
	}
	minLen := 999999
	for _, num := range m {
		if num.Freq != da {
			continue
		}
		if num.End-num.Start+1 < minLen {
			minLen = num.End - num.Start + 1
		}
	}
	return minLen
}

func main() {
	// nums := []int{1, 2, 2, 3, 1}
	// nums := []int{1, 2, 2, 3, 1, 4, 2}
	nums := []int{1}
	fmt.Println(findShortestSubArray(nums))
}
