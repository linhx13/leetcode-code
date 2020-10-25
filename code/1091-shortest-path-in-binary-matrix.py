from typing import List
from collections import deque


class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        if grid[0][0] != 0 or grid[m - 1][n - 1] != 0:
            return -1
        queue = deque()
        visited = set()
        queue.append((0, 0))
        visited.add((0, 0))
        dirs = [[0, 1], [0, -1], [-1, 0], [1, 0], [-1, -1], [-1, 1], [1, -1],
                [1, 1]]
        res = 1
        while queue:
            for _ in range(len(queue)):
                cur = queue.popleft()
                if cur[0] == m - 1 and cur[1] == n - 1:
                    return res
                for d in dirs:
                    r, c = cur[0] + d[0], cur[1] + d[1]
                    if r >= 0 and r < m and c >= 0 and c < n \
                       and grid[r][c] == 0:
                        k = (r, c)
                        if k not in visited:
                            queue.append(k)
                            visited.add(k)
            res += 1
        return -1
