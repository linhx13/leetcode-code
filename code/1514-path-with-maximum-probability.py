from typing import List
import collections
import math
import heapq


class Solution:
    def maxProbability(
        self,
        n: int,
        edges: List[List[int]],
        succProb: List[float],
        start: int,
        end: int,
    ) -> float:
        graph = [collections.defaultdict(dict) for _ in range(n)]
        for (a, b), p in zip(edges, succProb):
            graph[a][b] = graph[b][a] = -math.log(p)

        visited = [False for _ in range(n)]
        h = []
        res = 0.0
        heapq.heappush(h, (0, start))
        while len(h) > 0:
            t = heapq.heappop(h)
            if t[1] == end:
                res = math.exp(-t[0])
                break
            if not visited[t[1]]:
                visited[t[1]] = True
                for i, d in graph[t[1]].items():
                    if not visited[i]:
                        heapq.heappush(h, (t[0] + d, i))
        return res


if __name__ == "__main__":
    # n = 3
    # edges = [[0, 1], [1, 2], [0, 2]]
    # succProb = [0.5, 0.5, 0.2]
    # start = 0
    # end = 2
    # n = 3
    # edges = [[0, 1], [1, 2], [0, 2]]
    # succProb = [0.5, 0.5, 0.3]
    # start = 0
    # end = 2
    n = 3
    edges = [[0, 1]]
    succProb = [0.5]
    start = 0
    end = 2
    print(Solution().maxProbability(n, edges, succProb, start, end))
