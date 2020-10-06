package main

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := []int{}
	sum := 0
	for _, n := range nums {
		sum += n
		sums = append(sums, sum)
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if j >= len(this.sums) {
		j = len(this.sums) - 1
	}
	if i <= 0 {
		return this.sums[j]
	} else {
		return this.sums[j] - this.sums[i-1]
	}
}
