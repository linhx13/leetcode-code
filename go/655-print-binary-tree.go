package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	h := getHeight(root)
	w := (1 << h) - 1
	res := [][]string{}
	for i := 0; i < h; i++ {
		res = append(res, make([]string, w))
	}
	fill(root, res, 0, 0, w-1)
	return res
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func fill(root *TreeNode, res [][]string, h int, l int, r int) {
	if root == nil {
		return
	}
	mid := l + (r-l)>>1
	res[h][mid] = fmt.Sprintf("%d", root.Val)
	fill(root.Left, res, h+1, l, mid-1)
	fill(root.Right, res, h+1, mid+1, r)
}
