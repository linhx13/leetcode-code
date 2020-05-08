package main

import "fmt"

type Cashier struct {
	n         int
	discount  int
	prices    map[int]int
	customers int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	pp := make(map[int]int)
	for i := 0; i < len(products); i++ {
		pp[products[i]] = prices[i]
	}
	return Cashier{
		n:         n,
		discount:  discount,
		prices:    pp,
		customers: 0,
	}
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	bill := 0.0
	for i := 0; i < len(product); i++ {
		bill += float64(this.prices[product[i]] * amount[i])
	}
	this.customers += 1
	if this.customers%this.n == 0 {
		bill *= (1.0 - 1.0*float64(this.discount)/100)
	}
	return bill
}

func main() {
	cashier := Constructor(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	fmt.Println(cashier)
	fmt.Println(cashier.GetBill([]int{1, 2}, []int{1, 2}))
	fmt.Println(cashier.GetBill([]int{3, 7}, []int{10, 10}))
	fmt.Println(cashier.GetBill([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1}))
}
