package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	cache := make(map[int][]*TreeNode)
	res := helper(N, cache)
	return res
}

func helper(N int, cache map[int][]*TreeNode) []*TreeNode {
	if v, ok := cache[N]; ok {
		return v
	}
	res := []*TreeNode{}
	if N == 1 {
		root := &TreeNode{Val: 0}
		res = append(res, root)
	} else {
		for i := 1; i <= N-2; i++ {
			left := helper(i, cache)
			right := helper(N-1-i, cache)
			for _, l := range left {
				for _, r := range right {
					root := &TreeNode{Val: 0, Left: l, Right: r}
					res = append(res, root)
				}
			}
		}
	}
	cache[N] = res
	return res
}
