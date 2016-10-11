# use depth-first-search to build tree recursively

```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def dfs(self, head, length):
        if head is None or length == 0:
            return None
        if length == 1:
            root = TreeNode(head.val)
            return root
        left_length = length / 2
        right_length = length - left_length - 1
        left_head = head
        right_head = head
        i = 0
        mid = None
        while right_head is not None and i < left_length+1:
            if i == left_length:
                mid = right_head
            right_head = right_head.next
            i += 1
        root = TreeNode(mid.val)
        if left_head is None:
            root.left = None
        else:
            root.left = self.dfs(left_head, left_length)
        if right_head is None:
            root.right = None
        else:
            root.right = self.dfs(right_head, right_length)
        return root
    
    def sortedListToBST(self, head):
        """
        :type head: ListNode
        :rtype: TreeNode
        """
        if head is None:
            return None
        length = 0
        cur = head
        while cur is not None:
            cur = cur.next
            length += 1
        root = self.dfs(head, length)
        return root
        
```

