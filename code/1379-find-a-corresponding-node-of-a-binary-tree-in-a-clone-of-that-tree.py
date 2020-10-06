class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        if original is None:
            return None
        if original == target:
            return cloned
        if original.left is not None:
            res = self.getTargetCopy(original.left, cloned.left, target)
            if res is not None:
                return res
        if original.right is not None:
            res = self.getTargetCopy(original.right, cloned.right, target)
            if res is not None:
                return res
        return None
