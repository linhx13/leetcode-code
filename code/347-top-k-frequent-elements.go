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
	return h[i].freq > h[j].freq
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

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	h := &PairHeap{}
	heap.Init(h)
	for n, c := range m {
		heap.Push(h, Pair{num: n, freq: c})
	}
	res := []int{}
	for ; k > 0; k-- {
		p := heap.Pop(h).(Pair)
		res = append(res, p.num)
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
