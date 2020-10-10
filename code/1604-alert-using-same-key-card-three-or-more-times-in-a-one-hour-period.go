package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {
	m := make(map[string][]int)
	for i := 0; i < len(keyName); i++ {
		arr := strings.Split(keyTime[i], ":")
		hh, _ := strconv.Atoi(arr[0])
		mm, _ := strconv.Atoi(arr[1])
		m[keyName[i]] = append(m[keyName[i]], hh*60+mm)
	}
	res := []string{}
	for k, v := range m {
		sort.Ints(v)
		flag := false
		for i := 0; i < len(v); i++ {
			if i+1 < len(v) && v[i+1]-v[i] <= 60 && i+2 < len(v) && v[i+2]-v[i] <= 60 {
				flag = true
				break
			}
		}
		if flag {
			res = append(res, k)
		}
	}
	sort.Strings(res)
	return res
}

func main() {
	// keyName := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	// keyTime := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
	// keyName := []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"}
	// keyTime := []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"}
	// keyName := []string{"john", "john", "john"}
	// keyTime := []string{"23:58", "23:59", "00:01"}
	// keyName := []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"}
	// keyTime := []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"}
	keyName := []string{"a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "c", "c", "c", "c"}
	keyTime := []string{"01:35", "08:43", "20:49", "00:01", "17:44", "02:50", "18:48", "22:27", "14:12", "18:00", "12:38", "20:40", "03:59", "22:24"}
	fmt.Println(alertNames(keyName, keyTime)) //
}
