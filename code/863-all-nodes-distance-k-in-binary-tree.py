from typing import List, Dict
from collections import defaultdict, deque


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def distanceK(self, root: TreeNode, target: TreeNode, K: int) -> List[int]:
        graph = defaultdict(list)
        self.dfs(root, graph)
        res = self.bfs(graph, target, K)
        return res

    def dfs(self, root: TreeNode, graph: Dict):
        if root is None:
            return
        if root.left is not None:
            graph[root].append(root.left)
            graph[root.left].append(root)
            self.dfs(root.left, graph)
        if root.right is not None:
            graph[root].append(root.right)
            graph[root.right].append(root)
            self.dfs(root.right, graph)

    def bfs(self, graph: Dict, target: TreeNode, K: int) -> List[int]:
        res = []
        queue = deque()
        visited = set()
        queue.append(target)
        visited.add(target)
        dist = 0
        while queue:
            for _ in range(len(queue)):
                cur = queue.popleft()
                if dist == K:
                    res.append(cur.val)
                for v in graph[cur]:
                    if v not in visited:
                        queue.append(v)
                        visited.add(v)
            dist += 1
            if dist > K:
                break
        return res
