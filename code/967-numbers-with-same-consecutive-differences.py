from typing import List
from collections import deque


class Solution:
    def numsSameConsecDiff(self, n: int, k: int) -> List[int]:
        queue = deque()
        visited = set()
        res = []
        for i in range(10):
            queue.append(i)
            visited.add(i)
        while len(queue) > 0:
            cur = queue.popleft()
            if len(str(cur)) == n:
                res.append(cur)
                continue
            last = cur % 10
            if last - k >= 0:
                x = cur * 10 + (last - k)
                if x not in visited:
                    queue.append(x)
                    visited.add(x)
            if last + k <= 9:
                x = cur * 10 + (last + k)
                if x not in visited:
                    queue.append(x)
                    visited.add(x)
        return res


if __name__ == '__main__':
    n = 2
    k = 2
    print(Solution().numsSameConsecDiff(n, k))
