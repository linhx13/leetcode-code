package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(nums []int, left, right int) *TreeNode {
	if len(nums) == 0 || left >= right {
		return nil
	}
	max := -999999999
	maxIdx := -1
	for i := left; i < right; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}
	if maxIdx == -1 {
		return nil
	}
	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = helper(nums, left, maxIdx)
	root.Right = helper(nums, maxIdx+1, right)
	return root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return helper(nums, 0, len(nums))
}
