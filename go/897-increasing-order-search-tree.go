package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, last *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	var left_most *TreeNode = nil
	var right_most *TreeNode = nil
	if root.Left != nil {
		left_most, right_most = dfs(root.Left, last)
		right_most.Left = nil
		right_most.Right = root
	} else {
		left_most = root
	}
	if last != nil {
		last.Right = left_most
	}
	root.Left = nil
	if root.Right != nil {
		root.Right, right_most = dfs(root.Right, root)
	} else {
		right_most = root
	}
	return left_most, right_most
}

func increasingBST(root *TreeNode) *TreeNode {
	left_most, _ := dfs(root, nil)
	return left_most
}
