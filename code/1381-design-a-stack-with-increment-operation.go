package main

type CustomStack struct {
	maxSize int
	data    []int
	top     int
}

func Constructor(maxSize int) CustomStack {
	var s CustomStack
	s.maxSize = maxSize
	s.top = 0
	s.data = make([]int, maxSize, maxSize)
	return s
}

func (this *CustomStack) Push(x int) {
	if this.top < this.maxSize {
		this.data[this.top] = x
		this.top++
	}
}

func (this *CustomStack) Pop() int {
	if this.top == 0 {
		return -1
	} else {
		this.top--
		return this.data[this.top]
	}
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < this.top; i++ {
		this.data[i] += val
	}
}

func main() {

}
