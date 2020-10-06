package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	node := root
	for node != nil {
		next := node.Next
		if node.Child != nil {
			child := flatten(node.Child)
			node.Child = nil
			node.Next = child
			child.Prev = node
			for child.Next != nil {
				child = child.Next
			}
			child.Next = next
			if next != nil {
				next.Prev = child
			}
			node = child
		} else {
			node = node.Next
		}
	}
	return root
}
