package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	res, _ := helper(root)
	return res
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftRes, leftSum := helper(root.Left)
	rightRes, rightSum := helper(root.Right)
	diff := int(math.Abs(float64(leftSum - rightSum)))
	res := leftRes + rightRes + diff
	sum := leftSum + rightSum + root.Val
	return res, sum
}
