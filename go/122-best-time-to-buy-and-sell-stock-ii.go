package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	buyIdx := 0
	for buyIdx < len(prices) {
		for ; buyIdx < len(prices); buyIdx++ {
			if buyIdx+1 < len(prices) && prices[buyIdx] < prices[buyIdx+1] {
				break
			}
		}
		if buyIdx >= len(prices) {
			break
		}
		sellIdx := buyIdx + 1
		sellPrice := 0
		for ; sellIdx < len(prices); sellIdx++ {
			sellPrice = prices[sellIdx]
			if sellIdx+1 < len(prices) && prices[sellIdx+1] < prices[sellIdx] {
				break
			}
		}
		res += sellPrice - prices[buyIdx]
		buyIdx = sellIdx
	}
	return res
}

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	// prices := []int{1, 2, 3, 4, 5}
	prices := []int{1, 1, 2, 2, 3, 1}
	fmt.Println(maxProfit(prices))
}
