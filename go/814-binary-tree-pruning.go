package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode) bool {
	if root == nil {
		return false
	}
	left := helper(root.Left)
	right := helper(root.Right)
	if !left {
		root.Left = nil
	}
	if !right {
		root.Right = nil
	}
	if root.Val == 1 {
		return true
	}
	if left || right {
		return true
	} else {
		return false
	}
}

func pruneTree(root *TreeNode) *TreeNode {
	res := helper(root)
	if !res {
		return nil
	} else {
		return root
	}
}
