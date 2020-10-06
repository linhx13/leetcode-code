package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Item struct {
	name   string
	time   int
	amount int
	city   string
	idx    int
}

func invalidTransactions(transactions []string) []string {
	m := make(map[string][]Item)
	items := []Item{}
	for i := 0; i < len(transactions); i++ {
		arr := strings.Split(transactions[i], ",")
		name, city := arr[0], arr[3]
		time, _ := strconv.Atoi(arr[1])
		amount, _ := strconv.Atoi(arr[2])
		item := Item{name: name, time: time, amount: amount, city: city, idx: i}
		m[name] = append(m[name], item)
		items = append(items, item)
	}
	res := []string{}
	for _, item := range items {
		if item.amount > 1000 {
			res = append(res, transactions[item.idx])
			continue
		}
		found := false
		for _, other := range m[item.name] {
			if other.idx != item.idx && other.city != item.city && int(math.Abs(float64(item.time-other.time))) <= 60 {
				found = true
				break
			}
		}
		if found {
			res = append(res, transactions[item.idx])
		}
	}
	return res
}

func main() {
	transactions := []string{"alice,20,800,mtv", "alice,50,100,beijing"}
	fmt.Println(invalidTransactions(transactions))
}
