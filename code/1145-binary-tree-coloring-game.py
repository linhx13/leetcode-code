class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def btreeGameWinningMove(self, root: TreeNode, n: int, x: int) -> bool:
        if root is None:
            return False
        if root.val == x:
            left_cnt = self.count(root.left)
            right_cnt = self.count(root.right)
            parent_cnt = n - (left_cnt + right_cnt + 1)
            return (left_cnt > right_cnt + parent_cnt
                    or right_cnt > left_cnt + parent_cnt
                    or parent_cnt > left_cnt + right_cnt)
        return self.btreeGameWinningMove(root.left, n, x) \
            or self.btreeGameWinningMove(root.right, n, x)

    def count(self, root: TreeNode) -> int:
        if root is None:
            return 0
        return self.count(root.left) + self.count(root.right) + 1
