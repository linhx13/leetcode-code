package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ints := this.serializeInts(root, []int{})
	strs := []string{}
	for _, i := range ints {
		strs = append(strs, fmt.Sprintf("%d", i))
	}
	return strings.Join(strs, ",")
}

func (this *Codec) serializeInts(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = this.serializeInts(root.Left, res)
	res = this.serializeInts(root.Right, res)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	strs := strings.Split(data, ",")
	ints := []int{}
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	pos := 0

	INT_MAX := int(^uint(0) >> 1)
	INT_MIN := ^INT_MAX
	return this.deserializeInts(ints, &pos, INT_MIN, INT_MAX)
}

func (this *Codec) deserializeInts(ints []int, pos *int, min, max int) *TreeNode {
	if *pos >= len(ints) {
		return nil
	}
	if ints[*pos] < min || ints[*pos] > max {
		return nil
	}
	root := &TreeNode{Val: ints[*pos]}
	(*pos)++
	root.Left = this.deserializeInts(ints, pos, min, root.Val)
	root.Right = this.deserializeInts(ints, pos, root.Val, max)
	return root
}
