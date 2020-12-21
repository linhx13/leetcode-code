from typing import List


class Solution:
    def frogPosition(
        self, n: int, edges: List[List[int]], t: int, target: int
    ) -> float:
        graph = [set() for _ in range(n + 1)]
        for e in edges:
            graph[e[0]].add(e[1])
            graph[e[1]].add(e[0])

        visited = set()

        def dfs(cur_v, cur_t):
            if cur_t > t:
                return 0
            if cur_t == t and cur_v == target:
                return 1.0
            nodes = set(u for u in graph[cur_v] if u not in visited)
            if not nodes:
                return 1 if cur_v == target else 0
            res = 0
            for u in nodes:
                visited.add(u)
                res += dfs(u, cur_t + 1)
                visited.remove(u)
            res *= 1 / len(nodes)
            return res

        visited.add(1)
        return dfs(1, 0)


if __name__ == "__main__":
    # n = 7
    # edges = [[1, 2], [1, 3], [1, 7], [2, 4], [2, 6], [3, 5]]
    # t = 2
    # target = 4
    n = 7
    edges = [[1, 2], [1, 3], [1, 7], [2, 4], [2, 6], [3, 5]]
    t = 1
    target = 7
    # n = 7
    # edges = [[1, 2], [1, 3], [1, 7], [2, 4], [2, 6], [3, 5]]
    # t = 20
    # target = 6
    print(Solution().frogPosition(n, edges, t, target))
