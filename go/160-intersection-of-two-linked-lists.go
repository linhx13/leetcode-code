package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA, nodeB := headA, headB
	hitNil := false
	for {
		if nodeA == nodeB {
			return nodeA
		} else {
			nodeA = nodeA.Next
			if nodeA == nil {
				if hitNil {
					return nil
				}
				hitNil = true
				nodeA = headB
			}
			nodeB = nodeB.Next
			if nodeB == nil {
				nodeB = headA
			}
		}
	}
}
