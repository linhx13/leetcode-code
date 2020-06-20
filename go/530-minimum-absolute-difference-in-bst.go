package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res := int(^uint(0) >> 1)
	pre := -1
	inorder(root, &pre, &res)
	return res
}

func inorder(root *TreeNode, pre *int, res *int) {
	if root == nil {
		return
	}
	inorder(root.Left, pre, res)
	if *pre != -1 {
		if *res > root.Val-*pre {
			*res = root.Val - *pre
		}
	}
	*pre = root.Val
	inorder(root.Right, pre, res)
}
