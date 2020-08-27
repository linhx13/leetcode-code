package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	Node  *TreeNode
	Level int
}

func findBottomLeftValue(root *TreeNode) int {
	queue := []Pair{}
	level := 0
	res := 0
	queue = append(queue, Pair{Node: root, Level: 1})
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:len(queue)]
		if pair.Level > level {
			res = pair.Node.Val
			level = pair.Level
		}
		if pair.Node.Left != nil {
			queue = append(queue, Pair{Node: pair.Node.Left, Level: pair.Level + 1})
		}
		if pair.Node.Right != nil {
			queue = append(queue, Pair{Node: pair.Node.Right, Level: pair.Level + 1})
		}
	}
	return res
}
