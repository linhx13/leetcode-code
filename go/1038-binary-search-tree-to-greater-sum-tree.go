package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, largerSum int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}
	newRight, sumRight := dfs(root.Right, largerSum)
	newLeft, sumLeft := dfs(root.Left, root.Val+sumRight+largerSum)
	newRoot := TreeNode{Val: largerSum + sumRight + root.Val}
	newRoot.Left = newLeft
	newRoot.Right = newRight
	return &newRoot, root.Val + sumLeft + sumRight
}

func bstToGst(root *TreeNode) *TreeNode {
	newRoot, _ := dfs(root, 0)
	return newRoot
}
