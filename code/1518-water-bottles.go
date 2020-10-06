package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	res := 0
	res += numBottles
	for numBottles >= numExchange {
		newBottles := numBottles / numExchange
		res += newBottles
		numBottles = newBottles + numBottles%numExchange
	}
	return res
}

func main() {
	// numBottles, numExchange := 9, 3
	// numBottles, numExchange := 15, 4
	numBottles, numExchange := 1, 2
	fmt.Println(numWaterBottles(numBottles, numExchange))
}
