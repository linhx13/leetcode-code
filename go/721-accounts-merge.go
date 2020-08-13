package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	root := make(map[string]string)
	owner := make(map[string]string)
	m := make(map[string]map[string]bool)

	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			root[account[i]] = account[i]
			owner[account[i]] = account[0]
		}
	}
	for _, account := range accounts {
		p := find(account[1], root)
		for i := 2; i < len(account); i++ {
			root[find(account[i], root)] = p
		}
	}
	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			p := find(account[i], root)
			if _, ok := m[p]; !ok {
				m[p] = make(map[string]bool)
			}
			m[p][account[i]] = true
		}
	}

	res := [][]string{}
	for k, v := range m {
		emails := []string{}
		for x, _ := range v {
			emails = append(emails, x)
		}
		sort.Strings(emails)
		cur := []string{owner[k]}
		cur = append(cur, emails...)
		res = append(res, cur)
	}
	return res
}

func find(s string, root map[string]string) string {
	if root[s] != s {
		root[s] = find(root[s], root)
	}
	return root[s]
}
