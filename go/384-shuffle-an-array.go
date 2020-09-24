package main

import "math/rand"

type Solution struct {
	nums []int
	ori  []int
}

func Constructor(nums []int) Solution {
	ori := make([]int, len(nums))
	nnums := make([]int, len(nums))
	copy(ori, nums)
	copy(nnums, nums)
	return Solution{nums: nnums, ori: ori}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.ori
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := len(this.nums) - 1; i >= 0; i-- {
		idx := rand.Intn(i + 1)
		this.nums[i], this.nums[idx] = this.nums[idx], this.nums[i]
	}
	return this.nums
}

func main() {

}
