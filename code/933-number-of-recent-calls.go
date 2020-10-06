package main

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	for len(this.pings) > 0 {
		if this.pings[0] < t-3000 {
			this.pings = this.pings[1:]
		} else {
			break
		}
	}
	return len(this.pings)
}
