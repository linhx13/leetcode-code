package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, vals []int) []int {
	if root == nil {
		return vals
	}
	vals = inorder(root.Left, vals)
	vals = append(vals, root.Val)
	vals = inorder(root.Right, vals)
	return vals
}

func build(low, high int, vals []int) *TreeNode {
	if low > high {
		return nil
	}
	mid := low + (high-low)>>1
	root := TreeNode{Val: vals[mid]}
	root.Left = build(low, mid-1, vals)
	root.Right = build(mid+1, high, vals)
	return &root
}

func balanceBST(root *TreeNode) *TreeNode {
	vals := inorder(root, []int{})
	return build(0, len(vals)-1, vals)
}
