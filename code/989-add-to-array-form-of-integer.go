package main

import "fmt"

func reverse(A []int) []int {
	n := len(A)
	for i := 0; i < n/2; i++ {
		A[i], A[n-i-1] = A[n-i-1], A[i]
	}
	return A
}

func addToArrayForm(A []int, K int) []int {
	B := []int{}
	if K == 0 {
		B = append(B, 0)
	} else {
		for K > 0 {
			B = append(B, K%10)
			K = K / 10
		}
	}
	A = reverse(A)
	res := []int{}
	carry := 0
	n := len(A)
	if len(B) > n {
		n = len(B)
	}
	for i := 0; i < n || carry > 0; i++ {
		x := carry
		if i < len(A) {
			x += A[i]
		}
		if i < len(B) {
			x += B[i]
		}
		res = append(res, x%10)
		carry = x / 10
	}
	return reverse(res)
}

func main() {
	// A := []int{1, 2, 0, 0}
	// K := 34
	// A := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	// A := []int{0}
	K := 0
	// A := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	// K := 516
	fmt.Println(addToArrayForm(A, K))
}
