class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def __init__(self):
        self.last_node = None

    def flatten(self, root: TreeNode) -> None:
        if root is None:
            return None
        self.last_node = root
        ori_left, ori_right = root.left, root.right
        if ori_left:
            self.last_node.left = None
            self.last_node.right = ori_left
            self.flatten(ori_left)
        if ori_right:
            self.last_node.left = None
            self.last_node.right = ori_right
            self.flatten(ori_right)
