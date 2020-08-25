package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	res, _, _ := dfs(root)
	return res
}

func dfs(root *TreeNode) (res int, min int, max int) {
	if root == nil {
		return 0, 0, 0
	}
	leftRes, leftMin, leftMax := dfs(root.Left)
	rightRes, rightMin, rightMax := dfs(root.Right)
	min, max = root.Val, root.Val

	res = 0
	if root.Left != nil {
		if leftMin < min {
			min = leftMin
		}
		if leftMax > max {
			max = leftMax
		}
		x := int(math.Abs(float64(root.Val - leftMin)))
		if x > res {
			res = x
		}
		x = int(math.Abs(float64(root.Val - leftMax)))
		if x > res {
			res = x
		}
		if leftRes > res {
			res = leftRes
		}
	}
	if root.Right != nil {
		if rightMin < min {
			min = rightMin
		}
		if rightMax > max {
			max = rightMax
		}
		x := int(math.Abs(float64(root.Val - rightMin)))
		if x > res {
			res = x
		}
		x = int(math.Abs(float64(root.Val - rightMax)))
		if x > res {
			res = x
		}
		if rightRes > res {
			res = rightRes
		}
	}
	return res, min, max
}
