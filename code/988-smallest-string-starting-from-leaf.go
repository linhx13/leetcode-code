package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	res := ""
	for _, bytes := range dfs(root) {
		s := string(bytes)
		if res == "" {
			res = s
		} else if s < res {
			res = s
		}
	}
	return res
}

func dfs(root *TreeNode) [][]byte {
	if root == nil {
		return [][]byte{}
	}
	c := byte('a' + root.Val)
	if root.Left == nil && root.Right == nil {
		return [][]byte{[]byte{c}}
	}
	res := [][]byte{}
	if root.Left != nil {
		left := dfs(root.Left)
		for _, bytes := range left {
			bytes = append(bytes, c)
			res = append(res, bytes)
		}
	}
	if root.Right != nil {
		right := dfs(root.Right)
		for _, bytes := range right {
			bytes = append(bytes, c)
			res = append(res, bytes)
		}
	}
	return res
}
