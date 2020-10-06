package main

import "fmt"

func duplicateZeros(arr []int) {
	n := len(arr)
	zeros := 0
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			zeros++
		}
	}
	if zeros == 0 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		if i+zeros < n {
			arr[i+zeros] = arr[i]
		}
		if arr[i] == 0 {
			zeros--
			if i+zeros < n {
				arr[i+zeros] = 0
			}
		}
	}
}

func main() {
	arr := []int{8, 4, 5, 0, 0, 0, 0, 7}
	duplicateZeros(arr)
	fmt.Println(arr)
}
