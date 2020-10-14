package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	return helper(pre, 0, len(pre), post, 0, len(post))
}

func helper(pre []int, preLow, preHigh int, post []int, postLow, postHigh int) *TreeNode {
	if preLow == preHigh {
		return nil
	}
	root := &TreeNode{Val: pre[preLow]}
	if preLow+1 == preHigh {
		return root
	}
	idx := -1
	for i := postLow; i < postHigh; i++ {
		if post[i] == pre[preLow+1] {
			idx = i
			break
		}
	}
	var left *TreeNode = nil
	var right *TreeNode = nil
	leftLen := idx - postLow + 1
	if leftLen > 0 {
		left = helper(pre, preLow+1, preLow+1+leftLen, post, postLow, postLow+leftLen)
	}
	right = helper(pre, preLow+1+leftLen, preHigh, post, postLow+leftLen, postHigh-1)
	root.Left = left
	root.Right = right
	return root
}
