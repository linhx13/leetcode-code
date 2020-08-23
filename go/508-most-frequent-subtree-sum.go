package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	m := make(map[int]int)
	helper(root, m)
	items := [][]int{}
	for k, v := range m {
		items = append(items, []int{k, v})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i][1] > items[j][1]
	})
	res := []int{items[0][0]}
	for i := 1; i < len(items) && items[i][1] == items[0][1]; i++ {
		res = append(res, items[i][0])
	}
	return res
}

func helper(root *TreeNode, m map[int]int) int {
	if root == nil {
		return 0
	}
	leftSum := helper(root.Left, m)
	rightSum := helper(root.Right, m)
	sum := root.Val + leftSum + rightSum
	m[sum]++
	return sum
}
