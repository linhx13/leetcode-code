package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	ch   byte
	freq int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func frequencySort(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	h := &PairHeap{}
	heap.Init(h)
	for c, n := range m {
		heap.Push(h, Pair{ch: c, freq: n})
	}
	res := []byte{}
	for len(*h) > 0 {
		p := heap.Pop(h).(Pair)
		for i := 0; i < p.freq; i++ {
			res = append(res, p.ch)
		}
	}
	return string(res)
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
