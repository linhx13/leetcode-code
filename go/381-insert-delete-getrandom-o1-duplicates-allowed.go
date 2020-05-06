package main

import (
	"fmt"
	"math/rand"
)

type RandomizedCollection struct {
	vals  []int
	index map[int]map[int]bool
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	var r RandomizedCollection
	r.index = make(map[int]map[int]bool)
	return r
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	m, ok := this.index[val]
	if !ok {
		m = make(map[int]bool)
	}
	m[len(this.vals)] = true
	this.index[val] = m
	this.vals = append(this.vals, val)
	return len(this.index[val]) == 1
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	m, ok := this.index[val]
	if !ok || len(m) == 0 {
		return false
	}
	idx := -1
	for k, _ := range m {
		idx = k
		break
	}
	delete(m, idx)
	if len(this.vals)-1 != idx {
		t := this.vals[len(this.vals)-1]
		this.vals[idx] = t
		delete(this.index[t], len(this.vals)-1)
		this.index[t][idx] = true
	}
	this.vals = this.vals[:len(this.vals)-1]
	return true

}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

func main() {
	rc := Constructor()
	fmt.Println(rc)
	rc.Insert(10)
	fmt.Println(rc)
}
