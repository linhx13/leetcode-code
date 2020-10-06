package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	for i := 0; i < len(flowerbed) && cnt < n; i++ {
		f := flowerbed[i]
		if f != 0 {
			continue
		}
		var leftOk, rightOk bool
		if i == 0 {
			leftOk = true
		} else {
			leftOk = flowerbed[i-1] == 0
		}
		if i == len(flowerbed)-1 {
			rightOk = true
		} else {
			rightOk = flowerbed[i+1] == 0
		}
		if leftOk && rightOk {
			cnt++
			flowerbed[i] = 1
			i++
		}
	}
	return cnt == n
}
