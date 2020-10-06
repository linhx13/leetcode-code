package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	cache := make(map[*TreeNode][]int)
	max := 0
	dfs(root, cache, &max)
	return max
}

func dfs(root *TreeNode, cache map[*TreeNode][]int, max *int) {
	if root == nil {
		return
	}
	dfs(root.Left, cache, max)
	dfs(root.Right, cache, max)
	l := maxZigZag(root, 0, cache)
	r := maxZigZag(root, 1, cache)
	if l > *max {
		*max = l
	}
	if r > *max {
		*max = r
	}
}

func maxZigZag(root *TreeNode, dir int, cache map[*TreeNode][]int) int {
	if root == nil {
		return 0
	}
	if v, ok := cache[root]; ok {
		if v[dir] != 0 {
			return v[dir]
		}
	}
	res := 0
	if dir == 0 && root.Left != nil {
		res = maxZigZag(root.Left, 1, cache) + 1
	}
	if dir == 1 && root.Right != nil {
		res = maxZigZag(root.Right, 0, cache) + 1
	}
	if _, ok := cache[root]; !ok {
		cache[root] = make([]int, 2)
	}
	cache[root][dir] = res
	return res
}
