from typing import List
from collections import deque


class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int,
                          dst: int, K: int) -> int:
        g = [[] for _ in range(n)]
        for f in flights:
            g[f[0]].append((f[1], f[2]))
        q = deque()
        res = 999999999
        q.append((src, 0))
        while q and K >= -1:
            K -= 1
            for _ in range(len(q)):
                cur = q.popleft()
                if cur[0] == dst:
                    res = min(res, cur[1])
                    continue
                for e in g[cur[0]]:
                    if e[1] + cur[1] < res:
                        q.append((e[0], e[1] + cur[1]))
        if res < 999999999:
            return res
        else:
            return -1
