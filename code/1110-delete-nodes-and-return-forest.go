package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, dm map[int]bool, res []*TreeNode) []*TreeNode {
	if node == nil {
		return res
	}
	v := dm[node.Val]
	if !v {
		leftRes := dfs(node.Left, dm, []*TreeNode{})
		foundLeft := false
		for _, r := range leftRes {
			if r == node.Left {
				foundLeft = true
			} else {
				res = append(res, r)
			}
		}
		if !foundLeft {
			node.Left = nil
		}
		rightRes := dfs(node.Right, dm, []*TreeNode{})
		foundRight := false
		for _, r := range rightRes {
			if r == node.Right {
				foundRight = true
			} else {
				res = append(res, r)
			}
		}
		if !foundRight {
			node.Right = nil
		}
		res = append(res, node)
	} else {
		leftRes := dfs(node.Left, dm, []*TreeNode{})
		rightRes := dfs(node.Right, dm, []*TreeNode{})
		for _, r := range leftRes {
			res = append(res, r)
		}
		for _, r := range rightRes {
			res = append(res, r)
		}
	}
	return res
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	dm := make(map[int]bool)
	for _, n := range to_delete {
		dm[n] = true
	}
	return dfs(root, dm, []*TreeNode{})
}
