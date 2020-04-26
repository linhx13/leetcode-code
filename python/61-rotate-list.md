# use two pointers

```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def rotateRight(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if head is None:
            return None
        if k == 0:
            return head
        first = head
        l = 0
        while first is not None:
            first = first.next
            l += 1
        k %= l
        first = head
        second = head
        for i in xrange(k):
            first = first.next
            if first is None:
                first = head
        while first.next is not None:
            second = second.next
            first = first.next
        first.next = head
        head = second.next
        second.next = None
        return head
```

