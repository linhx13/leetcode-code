package main

import "fmt"

func findKthBit(n int, k int) byte {
	var res byte

	for n > 1 {
		if k*2 == 1<<n {
			res ^= 1
			break
		}
		if k*2 < 1<<n {
			n--
		} else {
			k = (1 << n) - k
			n--
			res ^= 1
		}
	}
	return res + '0'
}

func main() {
	n, k := 3, 1
	fmt.Println(findKthBit(n, k))
}
