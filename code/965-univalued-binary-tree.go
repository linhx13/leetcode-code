package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if val >= 0 && root.Val != val {
		return false
	}
	return dfs(root.Left, root.Val) && dfs(root.Right, root.Val)
}

func isUnivalTree(root *TreeNode) bool {
	return dfs(root, -1)
}
