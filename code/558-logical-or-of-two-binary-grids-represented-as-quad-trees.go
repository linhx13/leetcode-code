package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return quadTree1
		} else {
			return quadTree2
		}
	}
	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return quadTree2
		} else {
			return quadTree1
		}
	}
	res := &Node{}
	tl := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	tr := intersect(quadTree1.TopRight, quadTree2.TopRight)
	bl := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	br := intersect(quadTree1.BottomRight, quadTree2.BottomRight)

	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val {
		res.IsLeaf = true
		res.Val = tl.Val
	} else {
		res.TopLeft = tl
		res.TopRight = tr
		res.BottomLeft = bl
		res.BottomRight = br
	}
	return res
}
