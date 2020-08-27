package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	root = dfs(root, key)
	return root
}

func dfs(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil && root.Right != nil {
			root = root.Right
		} else if root.Right == nil && root.Left != nil {
			root = root.Left
		} else {
			node := root.Left
			for node != nil && node.Right != nil {
				node = node.Right
			}
			node.Right = root.Right
			root = root.Left
		}
	} else if key < root.Val {
		root.Left = dfs(root.Left, key)
	} else {
		root.Right = dfs(root.Right, key)
	}
	return root
}
