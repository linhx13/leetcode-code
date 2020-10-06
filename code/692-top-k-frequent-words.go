package main

import (
	"container/heap"
)

type Pair struct {
	word string
	freq int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	} else {
		return h[i].word < h[j].word
	}
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		c := m[word]
		m[word] = c + 1
	}
	h := &PairHeap{}
	heap.Init(h)
	for w, c := range m {
		heap.Push(h, Pair{word: w, freq: c})
	}
	res := []string{}
	for ; k > 0; k-- {
		p := heap.Pop(h).(Pair)
		res = append(res, p.word)
	}
	return res
}
