package main

func findSpecialInteger(arr []int) int {
	last := -1
	cnt := 0
	for _, n := range arr {
		if n != last {
			if last != -1 && cnt*4 > len(arr) {
				return last
			}
			cnt = 0
			last = -1
		}
		last = n
		cnt++
	}
	return last
}
