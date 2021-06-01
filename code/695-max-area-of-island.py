from typing import List


class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        visited = set()
        res = 0
        m, n = len(grid), len(grid[0])
        for i in range(m):
            for j in range(n):
                if (i, j) in visited:
                    continue
                area = [0]
                self.dfs(grid, visited, m, n, i, j, area)
                res = max(res, area[0])
        return res

    def dfs(self, grid, visited, m, n, i, j, area):
        if grid[i][j] == 0:
            return
        if (i, j) in visited:
            return
        visited.add((i, j))
        area[0] += 1
        if j - 1 >= 0:
            self.dfs(grid, visited, m, n, i, j - 1, area)
        if j + 1 < n:
            self.dfs(grid, visited, m, n, i, j + 1, area)
        if i - 1 >= 0:
            self.dfs(grid, visited, m, n, i - 1, j, area)
        if i + 1 < m:
            self.dfs(grid, visited, m, n, i + 1, j, area)
