package main

func containsPattern(arr []int, m int, k int) bool {
	for i := 0; i+(m*k) <= len(arr); i++ {
		cnt := 1
		pat := arr[i : i+m]
		for j := i + m; j < len(arr); j += m {
			if !equal(pat, arr[j:j+m]) {
				break
			}
			cnt++
			if cnt == k {
				return true
			}
		}
	}
	return false
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}
