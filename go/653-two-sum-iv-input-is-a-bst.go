package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(root *TreeNode, k int, m map[int]int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val] = 1
	return find(root.Left, k, m) || find(root.Right, k, m)
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int)
	return find(root, k, m)
}

func main() {

}
