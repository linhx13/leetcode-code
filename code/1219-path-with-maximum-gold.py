from typing import List


class Solution:
    def getMaximumGold(self, grid: List[List[int]]) -> int:
        res = [0, 0]
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] != 0:
                    self.helper(grid, i, j, res)
        return res[1]

    def helper(self, grid, x, y, res):
        if (x >= 0 and x < len(grid) and y >= 0 and y < len(grid[0])
                and grid[x][y] > 0 and grid[x][y] <= 100):
            res[0] += grid[x][y]
            grid[x][y] += 100
            res[1] = max(res[0], res[1])
            self.helper(grid, x + 1, y, res)
            self.helper(grid, x - 1, y, res)
            self.helper(grid, x, y - 1, res)
            self.helper(grid, x, y + 1, res)
            grid[x][y] -= 100
            res[0] -= grid[x][y]
