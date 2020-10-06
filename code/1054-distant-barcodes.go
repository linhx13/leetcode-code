package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	num  int
	freq int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].num < h[j].num
	} else {
		return h[i].freq > h[j].freq
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

func rearrangeBarcodes(barcodes []int) []int {
	m := make(map[int]int)
	for _, n := range barcodes {
		m[n]++
	}
	h := &PairHeap{}
	for n, c := range m {
		*h = append(*h, Pair{num: n, freq: c})
	}
	heap.Init(h)
	res := []int{}
	for len(*h) > 0 {
		p1 := heap.Pop(h).(Pair)
		res = append(res, p1.num)
		if len(*h) > 0 {
			p2 := heap.Pop(h).(Pair)
			res = append(res, p2.num)
			p2.freq--
			if p2.freq > 0 {
				heap.Push(h, p2)
			}
		}
		p1.freq--
		if p1.freq > 0 {
			heap.Push(h, p1)
		}
	}
	return res
}

func main() {
	// barcodes := []int{1, 1, 1, 2, 2, 2}
	// barcodes := []int{1, 1, 1, 1, 2, 2, 3, 3}
	// barcodes := []int{2, 2, 1, 3}
	barcodes := []int{1, 2, 1}
	fmt.Println(rearrangeBarcodes(barcodes))
}
