package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, path []*TreeNode, res *int) {
	if node == nil {
		return
	}
	if len(path) >= 2 && path[len(path)-2].Val%2 == 0 {
		*res += node.Val
	}
	path = append(path, node)
	dfs(node.Left, path, res)
	dfs(node.Right, path, res)
	path = path[:len(path)-1]
}

func sumEvenGrandparent(root *TreeNode) int {
	res := 0
	dfs(root, []*TreeNode{}, &res)
	return res
}
