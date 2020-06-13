package main

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	c := 0
	for candies > 0 {
		if candies < c+1 {
			res[c%num_people] += candies
		} else {
			res[c%num_people] += c + 1
		}
		c++
		candies -= c
	}
	return res
}
