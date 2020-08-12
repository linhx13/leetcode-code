package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = inorder(root.Left, res)
	res = append(res, root.Val)
	res = inorder(root.Right, res)
	return res
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	t1 := inorder(root1, []int{})
	t2 := inorder(root2, []int{})
	res := []int{}
	i, j := 0, 0
	for len(res) != len(t1)+len(t2) {
		if j >= len(t2) {
			res = append(res, t1[i])
			i++
		} else if i >= len(t1) {
			res = append(res, t2[j])
			j++
		} else {
			if t1[i] < t2[j] {
				res = append(res, t1[i])
				i++
			} else {
				res = append(res, t2[j])
				j++
			}
		}
	}
	return res
}
