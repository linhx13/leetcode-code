from typing import List


class Solution:
    def maximalNetworkRank(self, n: int, roads: List[List[int]]) -> int:
        graph = [set() for _ in range(n)]
        for x, y in roads:
            graph[x].add(y)
            graph[y].add(x)
        degrees = [0 for _ in range(n)]
        for i, s in enumerate(graph):
            degrees[i] = len(s)
        res = 0
        for x in range(n):
            for y in range(n):
                if x == y:
                    continue
                k = degrees[x] + degrees[y]
                if x in graph[y] or y in graph[x]:
                    k -= 1
                res = max(res, k)
        return res


if __name__ == "__main__":
    # n = 4
    # roads = [[0, 1], [0, 3], [1, 2], [1, 3]]
    # n = 5
    # roads = [[0, 1], [0, 3], [1, 2], [1, 3], [2, 3], [2, 4]]
    n = 8
    roads = [[0, 1], [1, 2], [2, 3], [2, 4], [5, 6], [5, 7]]
    sol = Solution()
    print(sol.maximalNetworkRank(n, roads))
