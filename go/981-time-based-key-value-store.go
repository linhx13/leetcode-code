package main

import "fmt"

type Item struct {
	Value     string
	Timestamp int
}

type TimeMap struct {
	m map[string][]Item
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]Item)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Item{Value: value, Timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	items, ok := this.m[key]
	if !ok {
		return ""
	}
	low, high := 0, len(items)-1
	fmt.Println(key, timestamp, items)
	for low <= high {
		mid := low + (high-low)/2
		if items[mid].Timestamp > timestamp {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if high < 0 {
		return ""
	} else {
		return items[high].Value
	}
}
