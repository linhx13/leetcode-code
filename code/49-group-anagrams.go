package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i int, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)
		m[key] = append(m[key], s)
	}
	res := [][]string{}
	for _, v := range m {
		sort.Strings(v)
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
