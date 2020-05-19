package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, res *int) (int, int) {
	if node == nil {
		return 0, 0
	}
	nLeftNodes, nLeftCoins := dfs(node.Left, res)
	nRightNodes, nRightCoins := dfs(node.Right, res)
	*res += int(math.Abs(float64(nLeftCoins) - float64(nLeftNodes)))
	*res += int(math.Abs(float64(nRightCoins) - float64(nRightNodes)))
	return nLeftNodes + nRightNodes + 1, nLeftCoins + nRightCoins + node.Val
}

func distributeCoins(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}
