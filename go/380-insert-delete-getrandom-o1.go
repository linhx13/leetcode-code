package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	vals  []int
	index map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	var rs RandomizedSet
	rs.index = make(map[int]int)
	return rs
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.index[val]
	if ok {
		return false
	}
	this.index[val] = len(this.vals)
	this.vals = append(this.vals, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.index[val]
	if !ok {
		return false
	}
	last := this.vals[len(this.vals)-1]
	this.index[last] = idx
	this.vals[idx] = last
	this.vals = this.vals[:len(this.vals)-1]
	delete(this.index, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	fmt.Println(this.vals)
	return this.vals[rand.Intn(len(this.vals))]
}
