package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	arr_a := strings.Split(a[0:len(a)-1], "+")
	a1, _ := strconv.Atoi(arr_a[0])
	a2, _ := strconv.Atoi(arr_a[1])
	arr_b := strings.Split(b[0:len(b)-1], "+")
	b1, _ := strconv.Atoi(arr_b[0])
	b2, _ := strconv.Atoi(arr_b[1])
	x := a1*b1 - a2*b2
	y := a1*b2 + a2*b1
	return fmt.Sprintf("%d+%di", x, y)
}

func main() {
	a := "1+-1i"
	b := "1+-1i"
	fmt.Println(complexNumberMultiply(a, b))
}
