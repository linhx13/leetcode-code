package main

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}
	diff := 0
	for i := 1; i < len(A); i++ {
		d := A[i] - A[i-1]
		if diff < 0 && d > 0 {
			return false
		}
		if diff > 0 && d < 0 {
			return false
		}
		if d != 0 {
			diff = d
		}
	}
	return true
}
