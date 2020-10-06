package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode, prev *int, res *int) {
	if node == nil {
		return
	}
	helper(node.Left, prev, res)
	if *prev != -1 {
		if node.Val-*prev < *res {
			*res = node.Val - *prev
		}
	}
	*prev = node.Val
	helper(node.Right, prev, res)
}

func minDiffInBST(root *TreeNode) int {
	res := int(^uint(0) >> 1)
	prev := -1
	helper(root, &prev, &res)
	return res
}
