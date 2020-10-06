package main

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	digitLogs := []string{}
	data := [][]string{}
	for _, log := range logs {
		idx := strings.Index(log, " ")
		if log[idx+1] >= '0' && log[idx+1] <= '9' {
			digitLogs = append(digitLogs, log)
			continue
		}
		data = append(data, []string{log[:idx], log[idx+1:]})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i][1] < data[j][1] || (data[i][1] == data[j][1] && data[i][0] < data[j][0])
	})
	res := []string{}
	for _, x := range data {
		res = append(res, x[0]+" "+x[1])
	}
	for _, log := range digitLogs {
		res = append(res, log)
	}
	return res
}
