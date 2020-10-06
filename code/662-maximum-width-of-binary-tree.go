package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	node *TreeNode
	idx  int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []Item{}
	queue = append(queue, Item{node: root, idx: 0})
	res := 0
	for len(queue) > 0 {
		cur := queue[len(queue)-1].idx - queue[0].idx + 1
		if cur > res {
			res = cur
		}
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[0]
			queue = queue[1:len(queue)]
			if item.node.Left != nil {
				queue = append(queue, Item{node: item.node.Left, idx: item.idx * 2})
			}
			if item.node.Right != nil {
				queue = append(queue, Item{node: item.node.Right, idx: item.idx*2 + 1})
			}
		}
	}
	return res
}
