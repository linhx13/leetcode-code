# use two pointers

```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def partition(self, head, x):
        """
        :type head: ListNode
        :type x: int
        :rtype: ListNode
        """
        if head is None:
            return None
        first_head = None
        first_tail = None
        second_head = None
        second_tail = None
        cur = head
        while cur is not None:
            if cur.val < x:
                if first_head is None:
                    first_head = cur
                else:
                    first_tail.next = cur
                first_tail = cur
            else:
                if second_head is None:
                    second_head = cur
                else:
                    second_tail.next = cur
                second_tail = cur
            cur = cur.next
        if second_tail is not None:
            second_tail.next = None
        if first_tail is not None:
            first_tail.next = second_head
            return first_head
        else:
            return second_head
```

