from typing import List
from collections import deque


class Solution:
    def shortestBridge(self, A: List[List[int]]) -> int:
        cur = 11
        m, n = len(A), len(A[0])
        for x in range(m):
            for y in range(n):
                if A[x][y] == 1:
                    self.dfs(A, m, n, x, y, cur)
                    cur += 1
        queue = deque()
        visited = set()
        for x in range(m):
            for y in range(n):
                if A[x][y] == 11:
                    visited.add((x, y))
                    queue.append((x, y))
        dd = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        res = 0
        while queue:
            sz = len(queue)
            while sz > 0:
                x, y = queue.popleft()
                for dx, dy in dd:
                    xx, yy = x + dx, y + dy
                    if 0 <= xx < m and 0 <= yy < n:
                        if A[xx][yy] == 12:
                            return res
                        elif (xx, yy) not in visited:
                            queue.append((xx, yy))
                            visited.add((xx, yy))
                sz -= 1
            res += 1
        return res

    def dfs(self, A, m, n, x, y, cur):
        if x < 0 or x >= m or y < 0 or y >= n:
            return
        if A[x][y] != 1:
            return
        A[x][y] = cur
        self.dfs(A, m, n, x - 1, y, cur)
        self.dfs(A, m, n, x + 1, y, cur)
        self.dfs(A, m, n, x, y - 1, cur)
        self.dfs(A, m, n, x, y + 1, cur)


if __name__ == "__main__":
    # A = [[0, 1], [1, 0]]
    A = [[0, 1, 0], [0, 0, 0], [0, 0, 1]]
    # A = [[1, 1, 1, 1, 1], [1, 0, 0, 0, 1], [1, 0, 1, 0, 1], [1, 0, 0, 0, 1],
    #      [1, 1, 1, 1, 1]]
    sol = Solution()
    print(sol.shortestBridge(A))
