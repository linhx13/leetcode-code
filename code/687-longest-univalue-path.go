package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, res)
	r := dfs(root.Right, res)
	pl, pr := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		pl = l + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		pr = r + 1
	}
	if *res < pl+pr {
		*res = pl + pr
	}
	if pl > pr {
		return pl
	} else {
		return pr
	}
}
