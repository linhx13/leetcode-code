package main

type MapSum struct {
	vals map[string]int
	sums map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{vals: make(map[string]int), sums: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	cur := val
	if v, ok := this.vals[key]; ok {
		cur -= v
	}
	this.vals[key] = val
	for i := 1; i <= len(key); i++ {
		this.sums[key[0:i]] += cur
	}
}

func (this *MapSum) Sum(prefix string) int {
	return this.sums[prefix]
}
