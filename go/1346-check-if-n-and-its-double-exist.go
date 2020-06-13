package main

func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	for _, n := range arr {
		if _, ok := m[n*2]; ok {
			return true
		}
		if n%2 == 0 {
			if _, ok := m[n/2]; ok {
				return true
			}
		}
		m[n] = true
	}
	return false
}
