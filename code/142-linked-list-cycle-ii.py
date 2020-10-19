from typing import List


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        if head is None:
            return None
        slow, fast = head, head
        if head.next is None:
            return None
        slow = head.next
        fast = head.next
        if fast.next is None:
            return None
        fast = fast.next
        while slow != fast:
            if slow.next is None:
                return None
            slow = slow.next
            if fast.next is None or fast.next.next is None:
                return None
            fast = fast.next.next
        nodes = set()
        node = slow.next
        nodes.add(node)
        while node != slow:
            node = node.next
            nodes.add(node)
        node = head
        while node not in nodes:
            node = node.next
        return node
