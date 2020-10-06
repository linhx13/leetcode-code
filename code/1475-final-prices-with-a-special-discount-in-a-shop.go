package main

func finalPrices(prices []int) []int {
	res := []int{}
	for i, p := range prices {
		j := i + 1
		for ; j < len(prices); j++ {
			if prices[j] <= p {
				break
			}
		}
		if j < len(prices) {
			p -= prices[j]
		}
		res = append(res, p)
	}
	return res
}
