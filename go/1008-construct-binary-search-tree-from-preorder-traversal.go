package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(preorder []int, left, right int) *TreeNode {
	if len(preorder) == 0 || left >= right {
		return nil
	}
	root := &TreeNode{Val: preorder[left]}
	idx := left + 1
	for ; idx < right; idx++ {
		if preorder[idx] > preorder[left] {
			break
		}
	}
	root.Left = helper(preorder, left+1, idx)
	root.Right = helper(preorder, idx, right)
	return root
}

func bstFromPreorder(preorder []int) *TreeNode {
	return helper(preorder, 0, len(preorder))
}
