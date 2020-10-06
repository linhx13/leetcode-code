package main

import "fmt"

func minimumSwap(s1 string, s2 string) int {
	m := make(map[string]int)

	res := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		k := string([]byte{s1[i], s2[i]})
		if _, ok := m[k]; ok {
			m[k]--
			if m[k] == 0 {
				delete(m, k)
			}
			res++
		} else {
			m[k]++
		}
	}
	if len(m)%2 != 0 {
		return -1
	}
	res += len(m)
	return res
}

func main() {
	// s1, s2 := "xx", "yy"
	// s1, s2 := "xx", "xy"
	s1, s2 := "xxyyxyxyxx", "xyyxyxxxyx"
	fmt.Println(minimumSwap(s1, s2))
}
