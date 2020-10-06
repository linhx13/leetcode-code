package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, L int, R int, sum *int) {
	if root == nil {
		return
	}
	if L <= root.Val && root.Val <= R {
		*sum += root.Val
	}
	if L < root.Val {
		dfs(root.Left, L, R, sum)
	}
	if root.Val < R {
		dfs(root.Right, L, R, sum)
	}
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	dfs(root, L, R, &sum)
	return sum
}
