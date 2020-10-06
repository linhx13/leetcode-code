package main

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if target >= letters[n-1] {
		return letters[0]
	}
	low, high := 0, n
	for low < high {
		mid := low + (high-low)/2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return letters[high]
}
