package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	xy2ns := make(map[string][]*TreeNode)
	x2ys := make(map[int]map[int]bool)
	dfs(root, 0, 0, xy2ns, x2ys)
	xs := []int{}
	for x, _ := range x2ys {
		xs = append(xs, x)
	}
	sort.Ints(xs)
	res := [][]int{}
	for _, x := range xs {
		ys := []int{}
		for y, _ := range x2ys[x] {
			ys = append(ys, y)
		}
		sort.Slice(ys, func(i, j int) bool {
			return ys[i] > ys[j]
		})
		cur := []int{}
		for _, y := range ys {
			xy := fmt.Sprintf("%d#%d", x, y)
			vs := []int{}
			for _, n := range xy2ns[xy] {
				vs = append(vs, n.Val)
			}
			sort.Ints(vs)
			cur = append(cur, vs...)
		}
		res = append(res, cur)
	}
	return res
}

func dfs(root *TreeNode, x, y int, xy2ns map[string][]*TreeNode, x2ys map[int]map[int]bool) {
	if root == nil {
		return
	}
	xy := fmt.Sprintf("%d#%d", x, y)
	xy2ns[xy] = append(xy2ns[xy], root)
	if _, ok := x2ys[x]; !ok {
		x2ys[x] = make(map[int]bool)
	}
	x2ys[x][y] = true
	if root.Left != nil {
		dfs(root.Left, x-1, y-1, xy2ns, x2ys)
	}
	if root.Right != nil {
		dfs(root.Right, x+1, y-1, xy2ns, x2ys)
	}
}
