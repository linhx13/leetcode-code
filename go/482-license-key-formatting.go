package main

import "strings"

func licenseKeyFormatting(S string, K int) string {
	s := strings.ToUpper(strings.Join(strings.Split(S, "-"), ""))
	res := []byte{}
	first := len(s) % K
	i := 0
	for ; i < first; i++ {
		res = append(res, s[i])
	}
	for i < len(s) {
		if len(res) > 0 {
			res = append(res, '-')
		}
		res = append(res, s[i:i+K]...)
		i += K
	}
	return string(res)
}
