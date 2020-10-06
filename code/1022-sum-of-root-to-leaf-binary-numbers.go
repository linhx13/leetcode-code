package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode, p int) int {
	if node == nil {
		return 0
	}
	cur := node.Val + p<<1
	if node.Left == nil && node.Right == nil {
		return cur
	}
	return helper(node.Left, cur) + helper(node.Right, cur)
}

func sumRootToLeaf(root *TreeNode) int {
	return helper(root, 0)
}

func main() {

}
