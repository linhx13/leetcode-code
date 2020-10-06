package main

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i := 0; i < len(indices); i++ {
		res[indices[i]] = s[i]
	}
	return string(res)
}
