package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	if root != nil {
		sum := 0
		dfs(root, &sum)
	}
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root.Right != nil {
		dfs(root.Right, sum)
	}
	tmp := root.Val
	root.Val += *sum
	*sum += tmp
	if root.Left != nil {
		dfs(root.Left, sum)
	}
}
