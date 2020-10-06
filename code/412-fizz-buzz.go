package main

import (
	"fmt"
	"strings"
)

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		var cur []string
		if i%3 == 0 {
			cur = append(cur, "Fizz")
		}
		if i%5 == 0 {
			cur = append(cur, "Buzz")
		}
		if cur == nil {
			cur = append(cur, fmt.Sprintf("%v", i))
		}
		res = append(res, strings.Join(cur, ""))
	}
	return res
}

func main() {
	fmt.Println(fizzBuzz(15))
}
