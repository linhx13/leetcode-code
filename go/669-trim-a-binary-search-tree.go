package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	newLeft := trimBST(root.Left, L, R)
	newRight := trimBST(root.Right, L, R)
	if L <= root.Val && root.Val <= R {
		root.Left = newLeft
		root.Right = newRight
		return root
	} else {
		if newLeft != nil {
			return newLeft
		} else if newRight != nil {
			return newRight
		} else {
			return nil
		}
	}
}
