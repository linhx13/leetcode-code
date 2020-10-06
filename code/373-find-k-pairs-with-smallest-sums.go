package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	x, y int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].x+h[i].y < h[j].x+h[j].y
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

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := [][]int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	h := &PairHeap{}
	heap.Init(h)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(h, Pair{x: nums1[i], y: nums2[j]})
		}
	}
	for len(res) < k && len(*h) > 0 {
		p := heap.Pop(h).(Pair)
		res = append(res, []int{p.x, p.y})
	}
	return res
}

func main() {
	// nums1 := []int{1, 7, 11}
	// nums2 := []int{2, 4, 6}
	// k := 3
	// nums1 := []int{1, 1, 2}
	// nums2 := []int{1, 2, 3}
	// k := 2
	nums1 := []int{1, 2}
	nums2 := []int{3}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
