package main

import "math/rand"

type Solution struct {
	rows int
	cols int
	last int
	m    map[int]int
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{rows: n_rows, cols: n_cols, last: n_rows*n_cols - 1, m: make(map[int]int)}
}

func (this *Solution) Flip() []int {
	x := rand.Intn(this.last + 1)
	r := x

	if v, ok := this.m[x]; ok {
		r = v
	}
	if x != this.last {
		if _, ok := this.m[this.last]; !ok {
			this.m[x] = this.last
		} else {
			this.m[x] = this.m[this.last]
		}
	}
	this.last--
	return []int{r / this.cols, r % this.cols}
}

func (this *Solution) Reset() {
	this.last = this.rows*this.cols - 1
	this.m = make(map[int]int)
}
