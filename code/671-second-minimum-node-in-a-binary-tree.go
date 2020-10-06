package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, min int) int {
	if root == nil {
		return -1
	}
	if root.Val > min {
		return root.Val
	}
	resLeft := dfs(root.Left, min)
	resRight := dfs(root.Right, min)

	if resLeft == -1 {
		return resRight
	}
	if resRight == -1 {
		return resLeft
	}
	if resLeft < resRight {
		return resLeft
	} else {
		return resRight
	}
}
