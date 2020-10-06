package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := 0, 0
	if root.Left != nil {
		leftDepth = getDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = getDepth(root.Right)
	}
	depth := 1 + leftDepth
	if 1+rightDepth > depth {
		depth = 1 + rightDepth
	}
	return depth
}

func getSum(root *TreeNode, curDepth int, depth int, sum *int) {
	if root == nil {
		return
	}
	if curDepth == depth {
		*sum += root.Val
	}
	if root.Left != nil {
		getSum(root.Left, curDepth+1, depth, sum)
	}
	if root.Right != nil {
		getSum(root.Right, curDepth+1, depth, sum)
	}
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := getDepth(root)
	sum := 0
	getSum(root, 1, depth, &sum)
	return sum
}

func main() {

}
