package main

import "fmt"

type CombinationIterator struct {
	chars string
	len   int
	idx   []int
	next  bool
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	iter := CombinationIterator{
		chars: characters, len: combinationLength,
		idx: make([]int, combinationLength)}
	for i := 0; i < combinationLength; i++ {
		iter.idx[i] = i
	}
	iter.next = true
	return iter
}

func (this *CombinationIterator) Next() string {
	res := []byte{}
	for i := 0; i < this.len; i++ {
		res = append(res, this.chars[this.idx[i]])
	}

	pos := -1
	for i := this.len - 1; i >= 0; i-- {
		if this.idx[i]+this.len-i < len(this.chars) {
			pos = i
			break
		}
	}
	if pos != -1 {
		this.idx[pos]++
		for i := pos + 1; i < this.len; i++ {
			this.idx[i] = this.idx[i-1] + 1
		}
	} else {
		this.next = false
	}
	return string(res)
}

func (this *CombinationIterator) HasNext() bool {
	return this.next
}

func main() {
	iter := Constructor("abcd", 3)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
