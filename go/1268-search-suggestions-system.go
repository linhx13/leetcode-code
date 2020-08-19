package main

import (
	"fmt"
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	m := make(map[string][]string)
	for _, p := range products {
		for i := 1; i <= len(p); i++ {
			prefix := p[0:i]
			m[prefix] = append(m[prefix], p)
		}
	}
	for k, v := range m {
		sort.Strings(v)
		if len(v) >= 3 {
			m[k] = v[0:3]
		}
	}
	res := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[0:i]
		v, ok := m[prefix]
		if !ok {
			v = []string{}
		}
		res = append(res, v)
	}
	return res
}

func main() {
	// products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	// searchWord := "mouse"
	// products := []string{"havana"}
	// searchWord := "havana"
	// products := []string{"bags", "baggage", "banner", "box", "cloths"}
	// searchWord := "bags"
	products := []string{"havana"}
	searchWord := "tatiana"
	fmt.Println(suggestedProducts(products, searchWord))
}
