package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	inorderRootIdex := -1
	for idx, v := range inorder {
		if v == rootVal {
			inorderRootIdex = idx
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	leftLen := inorderRootIdex
	root.Left = buildTree(preorder[1:1+leftLen], inorder[:leftLen])
	root.Right = buildTree(preorder[1+leftLen:], inorder[inorderRootIdex+1:])
	return root
}

func main() {

}
