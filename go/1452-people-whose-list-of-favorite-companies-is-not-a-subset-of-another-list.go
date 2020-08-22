package main

func peopleIndexes(favoriteCompanies [][]string) []int {
	index := make(map[string][]int)
	for idx, names := range favoriteCompanies {
		for _, name := range names {
			index[name] = append(index[name], idx)
		}
	}
	res := []int{}
	for idx, names := range favoriteCompanies {
		all := make(map[int]bool)
		for i, name := range names {
			if i == 0 {
				for _, x := range index[name] {
					all[x] = true
				}
			} else {
				all = intersection(all, index[name])
			}
		}
		if len(all) == 1 {
			res = append(res, idx)
		}
	}
	return res
}

func intersection(x map[int]bool, y []int) map[int]bool {
	res := make(map[int]bool)
	for _, i := range y {
		if x[i] {
			res[i] = true
		}
	}
	return res
}
