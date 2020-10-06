package main

import "fmt"

func largestTimeFromDigits(A []int) string {
	res := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				l := 6 - i - j - k
				hour := 10*A[i] + A[j]
				min := 10*A[k] + A[l]
				if hour < 24 && min < 60 {
					if res < hour*60+min {
						res = hour*60 + min
					}
				}
			}
		}
	}
	if res >= 0 {
		return fmt.Sprintf("%02d:%02d", res/60, res%60)
	} else {
		return ""
	}
}
