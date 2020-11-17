from typing import List
from collections import defaultdict


class Solution:
    def sumOfDistancesInTree(self, N: int,
                             edges: List[List[int]]) -> List[int]:
        graph = defaultdict(set)
        for e in edges:
            graph[e[0]].add(e[1])
            graph[e[1]].add(e[0])
        nodes = [1] * N
        dists = [0] * N

        def dfs(u, visited):
            if u in visited:
                return
            visited.add(u)
            for v in graph[u]:
                if v in visited:
                    continue
                dfs(v, visited)
                nodes[u] += nodes[v]
                dists[u] += nodes[v] + dists[v]

        dfs(0, set())

        res = [0] * N

        def dfs2(u, prev, visited):
            if u in visited:
                return
            visited.add(u)
            res[u] = dists[u]
            if prev != -1:
                res[u] += res[prev] - nodes[u] - dists[u] + N - nodes[u]
            for v in graph[u]:
                if v in visited:
                    continue
                dfs2(v, u, visited)

        dfs2(0, -1, set())
        return res


if __name__ == '__main__':
    N = 6
    edges = [[0, 1], [0, 2], [2, 3], [2, 4], [2, 5]]
    print(Solution().sumOfDistancesInTree(N, edges))
