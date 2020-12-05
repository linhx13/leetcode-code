class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxPathSum(self, root: TreeNode) -> int:
        res = [-100 * 3 * 10000]

        def dfs(node):
            if node is None:
                return 0
            left_res = dfs(node.left)
            right_res = dfs(node.right)
            res[0] = max(res[0], node.val, left_res + right_res + node.val,
                         left_res + node.val, right_res + node.val)
            return max(node.val, left_res + node.val, right_res + node.val)

        dfs(root)
        return res[0]
