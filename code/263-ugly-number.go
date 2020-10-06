package main

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	factors := []int{2, 3, 5}
	for _, f := range factors {
		for num != 0 && num%f == 0 {
			num = num / f
		}
	}
	return num == 1
}
