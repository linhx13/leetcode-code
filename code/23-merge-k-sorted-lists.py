from typing import List


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        k = len(lists)
        gap = 1
        while gap < k:
            for i in range(0, k - gap, gap * 2):
                lists[i] = self.merge2lists(lists[i], lists[i + gap])
            gap *= 2
        return lists[0] if k > 0 else None

    def merge2lists(self, l1, l2):
        head = ptr = ListNode(0)
        while l1 and l2:
            if l1.val <= l2.val:
                ptr.next = l1
                l1 = l1.next
            else:
                ptr.next = l2
                l2 = l2.next
            ptr = ptr.next
        if not l1:
            ptr.next = l2
        else:
            ptr.next = l1
        return head.next
