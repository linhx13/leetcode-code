package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cp := range cpdomains {
		arr := strings.Fields(cp)
		cnt, _ := strconv.Atoi(arr[0])
		parts := strings.Split(arr[1], ".")
		domains := make([]string, 0)
		for idx := len(parts) - 1; idx >= 0; idx-- {
			domains = append(domains, parts[idx])
			var new_domains []string
			for j := len(domains) - 1; j >= 0; j-- {
				new_domains = append(new_domains, domains[j])
			}
			d := strings.Join(new_domains, ".")
			c, ok := m[d]
			if ok {
				c += cnt
			} else {
				c = cnt
			}
			m[d] = c
		}
	}
	res := make([]string, 0)
	for k, v := range m {
		res = append(res, fmt.Sprintf("%v %v", v, k))
	}
	return res
}

func main() {
	// cpdomains := []string{"9001 discuss.leetcode.com"}
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(cpdomains))
}
