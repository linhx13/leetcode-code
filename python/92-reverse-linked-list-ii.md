# use pointer one-pass

```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseBetween(self, head, m, n):
        """
        :type head: ListNode
        :type m: int
        :type n: int
        :rtype: ListNode
        """
        if head is None:
            return head
        if m == n:
            return head
        prev = None
        i = 1
        while i < m:
            if prev == None:
                prev = head
            else:
                prev = prev.next
            i += 1
        if prev is None:
            cur = head
        else:
            cur = prev.next
        next = cur.next
        while i < n:
            if next is not None:
                next_next = next.next
                next.next = cur
                cur = next
                next = next_next
            i += 1
        if prev is not None:
            prev.next.next = next
            prev.next = cur
        else:
            head.next = next
            head = cur
        cur = head
        return head
```

