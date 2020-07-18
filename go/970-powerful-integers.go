package main

func powerfulIntegers(x int, y int, bound int) []int {
	res := make(map[int]bool)
	for i := 1; i <= bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			res[i+j] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	sliceRes := []int{}
	for k, _ := range res {
		sliceRes = append(sliceRes, k)
	}
	return sliceRes

}
