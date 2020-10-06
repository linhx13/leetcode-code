# tree traverse recursively

```python
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def do_preorder(self, root, ret_list):
        if root is None:
            return
        ret_list.append(root.val)
        if root.left is not None:
            self.do_preorder(root.left, ret_list)
        if root.right is not None:
            self.do_preorder(root.right, ret_list)
    
    def preorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        ret_list = []
        self.do_preorder(root, ret_list)
        return ret_list
```

# use stack

```python
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def preorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if root is None:
            return []
        ret_list = []
        stack = []
        stack.append([root, 0])
        while len(stack) > 0:
            if stack[-1][1] == 0:
                stack[-1][1] = 1
                ret_list.append(stack[-1][0].val)
            elif stack[-1][1] == 1:
                stack[-1][1] = 2
                if stack[-1][0].left is not None:
                    stack.append([stack[-1][0].left, 0])
            elif stack[-1][1] == 2:
                stack[-1][1] = 3
                if stack[-1][0].right is not None:
                    stack.append([stack[-1][0].right, 0])
            else:
                stack.pop()
        return ret_list
```

