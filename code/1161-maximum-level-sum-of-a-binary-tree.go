package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	node  *TreeNode
	level int
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := make(map[int]int)
	queue := []pair{}
	queue = append(queue, pair{node: root, level: 1})
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:len(queue)]
		m[p.level] += p.node.Val
		if p.node.Left != nil {
			queue = append(queue, pair{node: p.node.Left, level: p.level + 1})
		}
		if p.node.Right != nil {
			queue = append(queue, pair{node: p.node.Right, level: p.level + 1})
		}
	}
	max := ^int(^uint(0) >> 1)
	res := 0
	for k, v := range m {
		if v > max {
			res, max = k, v
		}
	}
	return res
}
