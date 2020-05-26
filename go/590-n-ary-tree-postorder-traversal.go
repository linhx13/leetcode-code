package main

type Node struct {
	Val      int
	Children []*Node
}

type Item struct {
	Node *Node
	idx  int
}

func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []Item{}
	stack = append(stack, Item{Node: root, idx: 0})
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		if item.idx >= len(item.Node.Children) {
			res = append(res, item.Node.Val)
			stack = stack[:len(stack)-1]
			continue
		}
		c := item.Node.Children[item.idx]
		item.idx++
		stack[len(stack)-1] = item
		stack = append(stack, Item{Node: c, idx: 0})
	}
	return res
}
