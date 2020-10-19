from typing import List
from collections import deque


class Solution:
    def validateBinaryTreeNodes(self, n: int, leftChild: List[int],
                                rightChild: List[int]) -> bool:
        parent = [-1 for _ in range(n)]
        for i in range(n):
            if leftChild[i] != -1:
                if leftChild[i] == i:
                    return False
                if parent[leftChild[i]] != -1:
                    return False
                parent[leftChild[i]] = i
            if rightChild[i] != -1:
                if rightChild[i] == i:
                    return False
                if parent[rightChild[i]] != -1:
                    return False
                parent[rightChild[i]] = i
        roots = [i for i in range(n) if parent[i] == -1]
        if len(roots) != 1:
            return False

        queue = deque()
        visited = [False for _ in range(n)]
        queue.append(roots[0])
        visited[roots[0]] = True
        while queue:
            cur = queue.popleft()
            l, r = leftChild[cur], rightChild[cur]
            if l >= 0:
                if visited[l]:
                    return False
                queue.append(l)
                visited[l] = True
            if r >= 0:
                if visited[r]:
                    return False
                queue.append(r)
                visited[r] = True
        return sum(visited) == n
