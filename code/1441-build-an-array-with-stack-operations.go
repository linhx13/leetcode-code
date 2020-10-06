package main

func buildArray(target []int, n int) []string {
	res := []string{}
	for idx, i := 0, 1; idx < len(target) && i <= n; {
		if target[idx] == i {
			res = append(res, "Push")
			idx++
			i++
		} else if target[idx] > i {
			res = append(res, "Push", "Pop")
			i++
		}
	}
	return res
}
