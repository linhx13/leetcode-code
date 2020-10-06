package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val == val:
		return root
	case val < root.Val:
		return searchBST(root.Left, val)
	case val > root.Val:
		return searchBST(root.Right, val)
	}
	return nil
}
