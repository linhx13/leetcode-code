package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode) (d int, l int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return 0, 0
	}
	ld, rd, cd := 0, 0, 0
	ll, rl := 0, 0
	if node.Left != nil {
		ld, ll = helper(node.Left)
		d += ll + 1
	}
	if node.Right != nil {
		rd, rl = helper(node.Right)
		d += rl + 1
	}
	if ld > rd {
		cd = ld
	} else {
		cd = rd
	}
	if cd > d {
		d = cd
	}
	if ll > rl {
		l = ll + 1
	} else {
		l = rl + 1
	}
	return d, l
}

func diameterOfBinaryTree(root *TreeNode) int {
	d, _ := helper(root)
	return d
}

func main() {

}
