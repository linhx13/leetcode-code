package main

func camelMatch(queries []string, pattern string) []bool {
	res := []bool{}
	for _, query := range queries {
		res = append(res, isMatch(query, pattern))
	}
	return res
}

func isMatch(query string, pattern string) bool {
	j := 0
	for i := 0; i < len(query); i++ {
		if j < len(pattern) && query[i] == pattern[j] {
			j++
		} else if query[i] < 'a' || query[i] > 'z' {
			return false
		}
	}
	return j == len(pattern)
}
