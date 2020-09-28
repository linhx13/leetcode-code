package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	cnts := make(map[int]int)
	res := 0
	dfs(root, cnts, &res)
	return res
}

func dfs(root *TreeNode, cnts map[int]int, res *int) {
	if root == nil {
		return
	}
	cnts[root.Val]++
	if root.Left == nil && root.Right == nil {
		c := 0
		for _, v := range cnts {
			if v%2 == 1 {
				c++
				if c > 1 {
					break
				}
			}
		}
		if c <= 1 {
			(*res)++
		}
	} else {
		dfs(root.Left, cnts, res)
		dfs(root.Right, cnts, res)
	}
	cnts[root.Val]--
}
