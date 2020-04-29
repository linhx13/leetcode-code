package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTargets(root *TreeNode, depth int, x int, y int) (px *TreeNode, py *TreeNode, dx int, dy int) {
	if root == nil {
		return nil, nil, 0, 0
	}
	if root.Left != nil {
		if root.Left.Val == x {
			px, dx = root, depth+1
		} else if root.Left.Val == y {
			py, dy = root, depth+1
		}
		px_, py_, dx_, dy_ := getTargets(root.Left, depth+1, x, y)
		if px_ != nil {
			px, dx = px_, dx_
		}
		if py_ != nil {
			py, dy = py_, dy_
		}
	}
	if root.Right != nil {
		if root.Right.Val == x {
			px, dx = root, depth+1
		} else if root.Right.Val == y {
			py, dy = root, depth+1
		}
		px_, py_, dx_, dy_ := getTargets(root.Right, depth+1, x, y)
		if px_ != nil {
			px, dx = px_, dx_
		}
		if py_ != nil {
			py, dy = py_, dy_
		}
	}
	return px, py, dx, dy
}

func isCousins(root *TreeNode, x int, y int) bool {
	px, py, dx, dy := getTargets(root, 0, x, y)
	return px != nil && py != nil && px != py && dx == dy
}

func main() {

}
