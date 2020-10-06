package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	m := 0
	for _, c := range root.Children {
		d := maxDepth(c)
		if d > m {
			m = d
		}
	}
	return m + 1
}
