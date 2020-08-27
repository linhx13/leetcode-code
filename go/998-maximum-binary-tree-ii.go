package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	return dfs(root, val)
}

func dfs(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		return &TreeNode{Val: val, Left: root}
	}
	root.Right = dfs(root.Right, val)
	return root
}
