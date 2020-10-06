package main

func generateTheString(n int) string {
	runes := []rune{}
	if n%2 == 1 {
		for i := 1; i <= n; i++ {
			runes = append(runes, 'a')
		}
	} else {
		for i := 1; i <= n-1; i++ {
			runes = append(runes, 'a')
		}
		runes = append(runes, 'b')
	}
	return string(runes)
}
