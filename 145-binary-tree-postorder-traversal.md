# use tree traverse recursively

```python
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def do_postorder(self, root, ret_list):
        if root is None:
            return
        if root.left is not None:
            self.do_postorder(root.left, ret_list)
        if root.right is not None:
            self.do_postorder(root.right, ret_list)
        ret_list.append(root.val)
        
    def postorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        ret_list = []
        self.do_postorder(root, ret_list)
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
    def do_postorder(self, root, ret_list):
        if root is None:
            return
        if root.left is not None:
            self.do_postorder(root.left, ret_list)
        if root.right is not None:
            self.do_postorder(root.right, ret_list)
        ret_list.append(root.val)
        
    def postorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        ret_list = []
        stack = []
        if root is not None:
            stack.append([root, 0])
        while len(stack) > 0:
            if stack[-1][1] == 0:
                stack[-1][1] = 1
                if stack[-1][0].left is not None:
                    stack.append([stack[-1][0].left, 0])
            elif stack[-1][1] == 1:
                stack[-1][1] = 2
                if stack[-1][0].right is not None:
                    stack.append([stack[-1][0].right, 0])
            elif stack[-1][1] == 2:
                stack[-1][1] = 3
                ret_list.append(stack[-1][0].val)
            else:
                stack.pop()
        return ret_list
```

