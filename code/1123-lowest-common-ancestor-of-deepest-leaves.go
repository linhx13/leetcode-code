package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode) (*TreeNode, int) {
	if node == nil {
		return nil, 0
	}
	leftRes, leftDepth := dfs(node.Left)
	rightRes, rightDepth := dfs(node.Right)
	if leftDepth > rightDepth {
		return leftRes, leftDepth + 1
	}
	if leftDepth < rightDepth {
		return rightRes, rightDepth + 1
	}
	return node, leftDepth + 1
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	res, _ := dfs(root)
	return res
}
