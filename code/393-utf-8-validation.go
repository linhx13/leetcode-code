package main

import "fmt"

func validUtf8(data []int) bool {
	cnt := 0
	for i := 0; i < len(data); i++ {
		val := data[i]
		if cnt != 0 {
			if val>>6 == 0b10 {
				cnt--
			} else {
				return false
			}
		} else {
			if val>>7 == 0b0 {
				cnt = 0
			} else if val>>5 == 0b110 {
				cnt = 1
			} else if val>>4 == 0b1110 {
				cnt = 2
			} else if val>>3 == 0b11110 {
				cnt = 3
			} else {
				return false
			}
		}
	}
	return cnt == 0
}

func main() {
	// data := []int{197, 130, 1}
	data := []int{235, 140, 4}
	fmt.Println(validUtf8(data))
}
