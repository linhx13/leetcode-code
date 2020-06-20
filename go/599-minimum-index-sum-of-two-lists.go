package main

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	for i, s := range list1 {
		m[s] = i
	}
	min := int(^uint(0) >> 1)
	res := []string{}
	for i, s := range list2 {
		c, ok := m[s]
		if ok {
			sum := i + c
			if sum == min {
				res = append(res, s)
			} else if sum < min {
				min = sum
				res = []string{s}
			}
		}
	}
	return res
}
