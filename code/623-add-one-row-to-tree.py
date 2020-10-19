class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def addOneRow(self, root: TreeNode, v: int, d: int) -> TreeNode:
        if d == 1:
            new_root = TreeNode(val=v, left=root)
            return new_root
        return self.dfs(root, v, d, 1)

    def dfs(self, root, v, d, depth):
        if root is None:
            return None
        if depth < d - 1:
            root.left = self.dfs(root.left, v, d, depth + 1)
            root.right = self.dfs(root.right, v, d, depth + 1)
        else:
            new_left = TreeNode(val=v, left=root.left)
            new_right = TreeNode(val=v, right=root.right)
            root.left = new_left
            root.right = new_right
        return root
