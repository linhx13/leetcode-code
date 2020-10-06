package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root *TreeNode
	m    map[int]*TreeNode
}

func Constructor(root *TreeNode) FindElements {
	if root == nil {
		return FindElements{root: root}
	}
	root.Val = 0
	m := make(map[int]*TreeNode)
	helper(root, m)
	return FindElements{root: root, m: m}
}

func helper(root *TreeNode, m map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
		m[root.Left.Val] = root.Left
	}
	helper(root.Left, m)
	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
		m[root.Right.Val] = root.Right
	}
	helper(root.Right, m)
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.m[target]
	return ok
}
