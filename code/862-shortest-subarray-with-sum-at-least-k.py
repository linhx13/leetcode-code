from typing import List
import collections


class Solution:
    def shortestSubarray(self, A: List[int], K: int) -> int:
        N = len(A)
        p = [0]
        for x in A:
            p.append(p[-1] + x)

        ans = N + 1
        queue = collections.deque()
        for y, py in enumerate(p):
            while queue and py <= p[queue[-1]]:
                queue.pop()

            while queue and py - p[queue[0]] >= K:
                ans = min(ans, y - queue.popleft())

            queue.append(y)
        return ans if ans < N + 1 else -1
