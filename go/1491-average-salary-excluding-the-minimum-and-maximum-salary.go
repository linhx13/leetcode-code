package main

func average(salary []int) float64 {
	min, max := int(^uint(0)>>1), 0
	sum := 0
	for _, x := range salary {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
		sum += x
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}
