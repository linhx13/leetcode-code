class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeInBetween(
        self, list1: ListNode, a: int, b: int, list2: ListNode
    ) -> ListNode:
        head = ListNode(next=list1)
        prev, cur = head, list1
        i = 0
        while i < a:
            prev = cur
            cur = cur.next
            i += 1
        while i < b:
            cur = cur.next
            i += 1
        next = cur.next
        prev.next = list2
        cur = list2
        while cur.next is not None:
            cur = cur.next
        cur.next = next
        return head.next
