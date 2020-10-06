package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getSum(root.Left) + getSum(root.Right) + root.Val
}

func findMaxProduct(root *TreeNode, sum int, max *int64) int {
	if root == nil {
		return 0
	}
	leftSum := findMaxProduct(root.Left, sum, max)
	rightSum := findMaxProduct(root.Right, sum, max)
	l := int64(leftSum * (sum - leftSum))
	if l > *max {
		*max = l
	}
	r := int64(rightSum * (sum - rightSum))
	if r > *max {
		*max = r
	}

	return leftSum + rightSum + root.Val
}

func maxProduct(root *TreeNode) int {
	sum := getSum(root)
	max := int64(0)
	findMaxProduct(root, sum, &max)
	return int(max % int64(1000000007))
}
