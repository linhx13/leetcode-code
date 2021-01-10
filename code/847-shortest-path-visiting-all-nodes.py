from typing import List
import collections


class Solution:
    def shortestPathLength(self, graph: List[List[int]]) -> int:
        N = len(graph)
        queue = collections.deque((1 << x, x) for x in range(N))
        dist = collections.defaultdict(lambda: N * N)
        for x in range(N):
            dist[(1 << x, x)] = 0

        while len(queue) > 0:
            cover, head = queue.popleft()
            d = dist[(cover, head)]
            if cover == 2 ** N - 1:
                return d
            for child in graph[head]:
                cover2 = cover | (1 << child)
                if d + 1 < dist[(cover2, child)]:
                    dist[(cover2, child)] = d + 1
                    queue.append((cover2, child))
