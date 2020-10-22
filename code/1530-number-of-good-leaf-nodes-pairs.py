class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


from collections import defaultdict, deque


class Solution:
    def countPairs(self, root: TreeNode, distance: int) -> int:
        graph = defaultdict(list)
        leaves = []
        self.build_graph(root, graph, leaves)
        res = 0
        for node in leaves:
            queue = deque()
            visited = set()
            queue.append(node)
            d = 0
            while queue:
                for _ in range(len(queue)):
                    cur = queue.popleft()
                    if d <= distance and cur is not node and cur.left is None and cur.right is None:
                        res += 1
                    for c in graph[cur]:
                        if c not in visited:
                            queue.append(c)
                            visited.add(c)
                d += 1
                if d > distance:
                    break
        return res // 2

    def build_graph(self, root: TreeNode, graph, leaves):
        if root is None:
            return
        if root.left is not None:
            graph[root].append(root.left)
            graph[root.left].append(root)
            self.build_graph(root.left, graph, leaves)
        if root.right is not None:
            graph[root].append(root.right)
            graph[root.right].append(root)
            self.build_graph(root.right, graph, leaves)
        if root.left is None and root.right is None:
            leaves.append(root)
