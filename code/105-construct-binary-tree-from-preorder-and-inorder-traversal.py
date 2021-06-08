from typing import List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        def helper(pl, pr, il, ir):
            if pl >= pr:
                return None
            if pl + 1 == pr:
                return TreeNode(val=preorder[pl])
            root = TreeNode(val=preorder[pl])
            root_idx = il
            for i in range(il, ir):
                if inorder[i] == preorder[pl]:
                    root_idx = i
                    break
            root.left = helper(pl + 1, pl + 1 + root_idx - il, il, root_idx)
            root.right = helper(pl + 1 + root_idx - il, pr, root_idx + 1, ir)
            return root

        return helper(0, len(preorder), 0, len(inorder))
