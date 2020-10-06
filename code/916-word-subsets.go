package main

import "fmt"

func wordSubsets(A []string, B []string) []string {
	mb := make(map[byte]int)
	cb := 0
	for _, w := range B {
		m := make(map[byte]int)
		for i := 0; i < len(w); i++ {
			m[w[i]]++
			cb |= (1 << (w[i] - 'a'))
		}
		for k, v := range m {
			if v > mb[k] {
				mb[k] = v
			}
		}
	}
	res := []string{}
	for _, w := range A {
		ma := make(map[byte]int)
		ca := 0
		for i := 0; i < len(w); i++ {
			ma[w[i]]++
			ca |= (1 << (w[i] - 'a'))
		}
		if cb&ca != cb {
			continue
		}
		valid := true
		for k, v := range mb {
			if v > ma[k] {
				valid = false
				break
			}
		}
		if valid {
			res = append(res, w)
		}
	}
	return res
}

func main() {
	// A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// B := []string{"e", "o"}
	// A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// B := []string{"l", "e"}
	// A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// B := []string{"e", "oo"}
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"ec", "oc", "ceo"}
	fmt.Println(wordSubsets(A, B))
}
