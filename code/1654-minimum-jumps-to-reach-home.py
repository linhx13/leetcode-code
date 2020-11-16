from typing import List
from collections import deque


class Solution:
    def minimumJumps(self, forbidden: List[int], a: int, b: int,
                     x: int) -> int:
        no_set = set(forbidden)
        queue = deque()
        visited = set()
        queue.append((0, 0))
        visited.add((0, 0))
        res = 0
        max_val = max([x] + forbidden) + a + b
        while queue:
            for _ in range(len(queue)):
                cur = queue.popleft()
                print(cur)
                if cur[0] == x:
                    return res
                t = cur[0] + a
                if t <= max_val and t not in no_set:
                    st = (t, 0)
                    if st not in visited:
                        visited.add(st)
                        queue.append(st)
                t = cur[0] - b
                if t >= 0 and cur[1] == 0 and t not in no_set:
                    st = (t, 1)
                    if st not in visited:
                        visited.add(st)
                        queue.append(st)
            res += 1
        return -1
