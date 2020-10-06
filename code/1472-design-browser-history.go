package main

type BrowserHistory struct {
	stack []string
	cur   int
	top   int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{stack: []string{homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	this.stack = this.stack[0 : this.cur+1]
	this.stack = append(this.stack, url)
	this.cur++
	this.top = this.cur
}

func (this *BrowserHistory) Back(steps int) string {
	this.cur -= steps
	if this.cur < 0 {
		this.cur = 0
	}
	return this.stack[this.cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.cur += steps
	if this.cur > this.top {
		this.cur = this.top
	}
	return this.stack[this.cur]
}
