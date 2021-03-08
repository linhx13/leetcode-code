from typing import List
import heapq


class Solution:
    def eatenApples(self, apples: List[int], days: List[int]) -> int:
        q = []
        i = 0
        res = 0
        while i < len(apples) or q:
            if i < len(apples) and apples[i] > 0:
                heapq.heappush(q, [i + days[i], apples[i]])
            while q and (q[0][0] <= i or q[0][1] == 0):
                heapq.heappop(q)
            if q:
                q[0][1] -= 1
                res += 1
            i += 1
        return res
