package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	if root == nil {
		return []int{-1}
	}
	idx := 0
	res := []int{}
	flag := true
	helper(root, voyage, &idx, &res, &flag)
	if !flag {
		res = []int{-1}
	}
	return res
}

func helper(root *TreeNode, voyage []int, idx *int, res *[]int, flag *bool) {
	if root == nil || *idx >= len(voyage) {
		return
	}
	x := voyage[*idx]
	(*idx)++
	if root.Val != x {
		*flag = false
		return
	}
	if root.Left != nil && root.Left.Val != voyage[*idx] {
		*res = append(*res, root.Val)
		root.Left, root.Right = root.Right, root.Left
	}
	helper(root.Left, voyage, idx, res, flag)
	if *flag == false {
		return
	}
	helper(root.Right, voyage, idx, res, flag)
	if *flag == false {
		return
	}
}
