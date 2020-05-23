package main

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	res := [][]int{}
	for idx, size := range groupSizes {
		group := m[size]
		group = append(group, idx)
		m[size] = group
		if len(group) == size {
			res = append(res, group)
			delete(m, size)
		}
	}
	return res
}
