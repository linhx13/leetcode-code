package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	maxDepth := getDepth(root, 1)
	return dfs(root, 1, maxDepth)
}

func getDepth(root *TreeNode, depth int) int {
	if root == nil {
		return 0
	}
	ld := getDepth(root.Left, depth+1)
	rd := getDepth(root.Right, depth+1)
	if ld == 0 && rd == 0 {
		return depth
	}
	if ld > rd {
		return ld
	} else {
		return rd
	}
}

func dfs(root *TreeNode, depth int, maxDepth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == maxDepth {
		return root
	}
	left := dfs(root.Left, depth+1, maxDepth)
	right := dfs(root.Right, depth+1, maxDepth)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}

func main() {

}
