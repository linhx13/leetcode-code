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

func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []Pair{}
	INT_MAX := int(^uint(0) >> 1)
	INT_MIN := ^INT_MAX
	level, max := 0, INT_MIN
	queue = append(queue, Pair{Node: root, Level: 1})
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:len(queue)]
		if level < pair.Level {
			if max != INT_MIN {
				res = append(res, max)
			}
			level = pair.Level
			max = INT_MIN
		}
		if pair.Node.Val > max {
			max = pair.Node.Val
		}
		if pair.Node.Left != nil {
			queue = append(queue, Pair{Node: pair.Node.Left, Level: pair.Level + 1})
		}
		if pair.Node.Right != nil {
			queue = append(queue, Pair{Node: pair.Node.Right, Level: pair.Level + 1})
		}
	}
	res = append(res, max)
	return res
}
