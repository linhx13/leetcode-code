package main

import "fmt"

func superPow(a int, b []int) int {
	mod := 1337
	a %= mod

	res := 1
	for i := 0; i < len(b); i++ {
		for j, f := b[i], a; j != 0; f, j = (f*f)%mod, j/2 {
			if j%2 == 1 {
				res = (res * f) % mod
			}
		}
		if i < len(b)-1 {
			for j, f := 9, res; j != 0; f, j = (f*f)%mod, j/2 {
				if j%2 == 1 {
					res = (res * f) % mod
				}
			}
		}
	}
	return res % mod
}

func main() {
	a := 2147483647
	b := []int{2, 0, 0}
	fmt.Println(superPow(a, b))
}
