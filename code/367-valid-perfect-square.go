package main

func isPerfectSquare(num int) bool {
	left, right := 0, num+1
	for left < right {
		mid := left + (right-left)/2
		s := mid * mid
		if s == num {
			return true
		} else if s > num {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}
