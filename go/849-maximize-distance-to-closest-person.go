package main

func maxDistToClosest(seats []int) int {
	front := make([]int, len(seats))
	back := make([]int, len(seats))
	for i := 1; i < len(seats); i++ {
		if seats[i] == 0 {
			front[i] = front[i-1] + 1
		}
	}
	for i := len(seats) - 2; i >= 0; i-- {
		if seats[i] == 0 {
			back[i] = back[i+1] + 1
		}
	}
	res := 0
	for i := 0; i < len(seats); i++ {
		cur := front[i]
		if back[i] < cur {
			cur = back[i]
		}
		if cur > res {
			res = cur
		}
	}
	if front[len(seats)-1] > res {
		res = front[len(seats)-1]
	}
	if back[0] > res {
		res = back[0]
	}
	return res
}
