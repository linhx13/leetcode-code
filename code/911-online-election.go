package main

type TopVotedCandidate struct {
	votes map[int]int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	votes := make(map[int]int)
	cnts := make(map[int]int)
	max, person := 0, -1
	for i := 0; i < len(persons); i++ {
		cnts[persons[i]]++
		if cnts[persons[i]] >= max {
			max = cnts[persons[i]]
			person = persons[i]
		}
		votes[times[i]] = person
	}
	return TopVotedCandidate{votes: votes, times: times}
}

func (this *TopVotedCandidate) Q(t int) int {
	key := this.floorKey(t)
	return this.votes[key]
}

func (this *TopVotedCandidate) floorKey(t int) int {
	low, high := 0, len(this.times)
	for low < high {
		mid := low + (high-low)>>1
		if this.times[mid] == t {
			return t
		} else if this.times[mid] > t {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return this.times[low-1]
}
