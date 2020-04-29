package main

import "fmt"

func shiftingLetters(S string, shifts []int) string {
	bytes := []byte(S)
	a := []byte("a")[0]
	sh := 0
	for i := len(shifts) - 1; i >= 0; i-- {
		sh += shifts[i]
		cur := (byte(sh%26) + bytes[i] - a) % 26
		bytes[i] = a + cur
	}
	return string(bytes)
}

func main() {
	fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
}
