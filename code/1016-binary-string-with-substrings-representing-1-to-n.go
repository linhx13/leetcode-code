package main

import (
	"fmt"
	"strconv"
	"strings"
)

func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		ns := strconv.FormatInt(int64(i), 2)
		if !strings.Contains(S, ns) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(queryString("110101011011000011011111000000", 15))
}
