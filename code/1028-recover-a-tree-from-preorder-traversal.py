class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def recoverFromPreorder(self, S: str) -> TreeNode:
        return self.helper(S, 0, -1)[0]

    def helper(self, S: str, idx: int, last_depth: int):
        node, depth, new_idx = self.parse(S, idx)
        if node is None:
            return None, idx
        if last_depth + 1 != depth:
            return None, idx
        idx = new_idx
        left_node, left_idx = self.helper(S, idx, depth)
        if left_node is None:
            return node, idx
        node.left = left_node
        right_node, right_idx = self.helper(S, left_idx, depth)
        if right_node is None:
            return node, left_idx
        node.right = right_node
        return node, right_idx

    def parse(self, S: str, idx: int):
        if not S:
            return None, -1, idx
        depth = 0
        while S[idx] == '-':
            depth += 1
            idx += 1
        k = S.find('-', idx)
        if k == -1:
            val = int(S[idx:])
        else:
            val = int(S[idx:k])
        node = TreeNode(val)
        return TreeNode(val), depth, k


def dump_tree(node):
    res = []
    from collections import deque
    queue = deque()
    queue.append(node)
    while queue:
        cur = queue.popleft()
        if cur:
            res.append(cur.val)
            if cur.left or cur.right:
                queue.append(cur.left)
                queue.append(cur.right)
        else:
            res.append(cur)
    return res


if __name__ == "__main__":
    # S = "1-2--3--4-5--6--7"
    S = "1-2--3---4-5--6---7"
    node = Solution().recoverFromPreorder(S)
    print(dump_tree(node))
