from typing import List


class Solution:
    def uniquePathsIII(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        res = [0]

        def dfs(i, j, empty):
            if not (0 <= i < m and 0 <= j < n and grid[i][j] >= 0):
                return
            if grid[i][j] == 2:
                if empty == 0:
                    res[0] += 1
                return
            grid[i][j] = -2
            dfs(i - 1, j, empty - 1)
            dfs(i + 1, j, empty - 1)
            dfs(i, j - 1, empty - 1)
            dfs(i, j + 1, empty - 1)
            grid[i][j] = 0

        empty = 1
        r, c = -1, -1
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    r, c = i, j
                elif grid[i][j] == 0:
                    empty += 1

        res = [0]
        dfs(r, c, empty)
        return res[0]
