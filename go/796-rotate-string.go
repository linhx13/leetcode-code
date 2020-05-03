package main

import (
	"fmt"
	"strings"
)

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	A = A + A
	return strings.Contains(A, B)
}

func main() {
	A := "1"
	B := ""
	fmt.Println(rotateString(A, B))
}
