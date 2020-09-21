package main

import "fmt"

func partitionLabels(S string) []int {
	res := []int{}
	if len(S) == 0 {
		return res
	}
	indeices := make([]int, 26)
	for i := 0; i < len(S); i++ {
		indeices[S[i]-'a'] = i
	}

	start, end := 0, 0
	for i := 0; i < len(S); i++ {
		end = max(end, indeices[S[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	S := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(S))
}
