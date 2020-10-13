package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{}
	level := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		level++
		firstNo, lastYes := -1, -1
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:len(queue)]
			if cur.Left != nil {
				lastYes = 2 * i
				queue = append(queue, cur.Left)
			} else {
				if firstNo == -1 {
					firstNo = 2 * i
				}
			}
			if cur.Right != nil {
				lastYes = 2*i + 1
				queue = append(queue, cur.Right)
			} else {
				if firstNo == -1 {
					firstNo = 2*i + 1
				}
			}
		}
		if firstNo != -1 && firstNo < lastYes {
			return false
		}
		if lastYes != -1 && sz != (1<<(level-1)) {
			return false
		}
	}
	return true
}
