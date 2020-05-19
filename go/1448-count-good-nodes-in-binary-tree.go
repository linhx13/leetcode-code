package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, max int, res *int) {
	if node == nil {
		return
	}
	if node.Val >= max {
		*res++
		max = node.Val
	}
	dfs(node.Left, max, res)
	dfs(node.Right, max, res)
}

func goodNodes(root *TreeNode) int {
	res := 0
	dfs(root, -99999, &res)
	return res
}
