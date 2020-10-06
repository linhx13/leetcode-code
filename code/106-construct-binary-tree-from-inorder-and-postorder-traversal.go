package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	inorderRootIdx := -1
	for idx, v := range inorder {
		if v == rootVal {
			inorderRootIdx = idx
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	leftLen := inorderRootIdx
	rightLen := len(inorder) - inorderRootIdx - 1
	root.Left = buildTree(inorder[:leftLen], postorder[:leftLen])
	root.Right = buildTree(inorder[inorderRootIdx+1:], postorder[leftLen:leftLen+rightLen])
	return root
}

func main() {

}
