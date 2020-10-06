package main

import "sort"

func numSpecialEquivGroups(A []string) int {
	m := make(map[string]bool)
	for _, s := range A {
		odd := []rune{}
		even := []rune{}
		for i, r := range s {
			if i%2 == 0 {
				even = append(even, r)
			} else {
				odd = append(odd, r)
			}
		}
		sort.Slice(even, func(i, j int) bool {
			return even[i] < even[j]
		})
		sort.Slice(odd, func(i, j int) bool {
			return odd[i] < odd[j]
		})
		odd = append(odd, even...)
		m[string(odd)] = true
	}
	return len(m)
}
