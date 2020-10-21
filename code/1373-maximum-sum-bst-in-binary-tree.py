class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxSumBST(self, root: TreeNode) -> int:
        res = [-1e9]
        self.dfs(root, res)
        return max(0, res[0])

    def dfs(self, root: TreeNode, res):
        if root is None:
            return (0, True, -1e9, 1e9)
        left = self.dfs(root.left, res)
        right = self.dfs(root.right, res)
        cur = left[0] + right[0] + root.val
        if left[1] and right[1] and left[2] < root.val and right[3] > root.val:
            res[0] = max(res[0], cur)
            return (cur, True, max(root.val, right[2]), min(root.val, left[3]))
        return (cur, False, max(root.val, right[2]), min(root.val, left[3]))
