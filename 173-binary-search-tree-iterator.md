# python code

```python
# Definition for a  binary tree node
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class BSTIterator(object):
    def __init__(self, root):
        """
        :type root: TreeNode
        """
        self.stack = [] 
        cur = root
        while cur is not None:
            self.stack.append([cur, 0])
            cur = cur.left
        if len(self.stack) > 0:
            self.stack[-1][1] = 1
        self.accessed_next = False
        
    def _move_to_next(self):
        while len(self.stack) > 0:
            if self.stack[-1][1] == 0:
                # has traversed left child
                self.stack[-1][1] = 1
                return
            elif self.stack[-1][1] == 1:
                # has traversed current node
                self.stack[-1][1] = 2
                cur = self.stack[-1][0].right
                if cur is None:
                    self.stack.pop()
                    continue
                while cur is not None:
                    self.stack.append(([cur, 0]))
                    cur = cur.left
                self.stack[-1][1] = 1
                return
            elif self.stack[-1][1] == 2:
                # has traversed right child
                self.stack.pop()
        

    def hasNext(self):
        """
        :rtype: bool
        """
        if self.accessed_next:
            self._move_to_next()
        self.accessed_next = False
        return len(self.stack) > 0

    def next(self):
        """
        :rtype: int
        """
        if self.accessed_next:
            self._move_to_next()
        self.accessed_next = True
        if len(self.stack) > 0:
            return self.stack[-1][0].val
        else:
            return None

# Your BSTIterator will be called like this:
# i, v = BSTIterator(root), []
# while i.hasNext(): v.append(i.next())
```

