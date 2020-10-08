package main

import "sort"

type MyCalendarTwo struct {
	m map[int]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{m: make(map[int]int)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	this.m[start]++
	this.m[end]--

	keys := []int{}
	for k, _ := range this.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	cnt := 0
	for _, k := range keys {
		cnt += this.m[k]
		if cnt > 2 {
			this.m[start]--
			this.m[end]++

			if this.m[start] == 0 {
				delete(this.m, start)
			}
			if this.m[end] == 0 {
				delete(this.m, end)
			}
			return false
		}
	}
	return true
}
