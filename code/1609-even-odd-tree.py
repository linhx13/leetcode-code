from collections import deque


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isEvenOddTree(self, root: TreeNode) -> bool:
        queue = deque()
        queue.append((root, 0))
        while queue:
            last = -1
            for _ in range(len(queue)):
                node, level = queue.popleft()
                if level & 1 == 0:
                    if node.val & 1 != 1:
                        return False
                    if last != -1 and not node.val > last:
                        return False
                elif level & 1 == 1:
                    if node.val & 1 != 0:
                        return False
                    if last != -1 and not node.val < last:
                        return False
                last = node.val
                if node.left is not None:
                    queue.append((node.left, level + 1))
                if node.right is not None:
                    queue.append((node.right, level + 1))
        return True
