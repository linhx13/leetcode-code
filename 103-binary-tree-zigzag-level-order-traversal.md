# use stack and breadth-first search

```python
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def zigzagLevelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        import collections
        queue = collections.deque()
        ret_list = []
        if root is not None:
            queue.append(root)
        stack = []
        is_reverse = False
        while len(queue) > 0:
            this_row = []
            while len(queue) > 0:
                node = queue.popleft()
                this_row.append(node.val)
                if is_reverse:
                    if node.right is not None:
                        stack.append(node.right)
                    if node.left is not None:
                        stack.append(node.left)
                else:
                    if node.left is not None:
                        stack.append(node.left)
                    if node.right is not None:
                        stack.append(node.right)
            ret_list.append(this_row)
            while len(stack) > 0:
                queue.append(stack.pop())
            is_reverse = not is_reverse
        return ret_list
```

