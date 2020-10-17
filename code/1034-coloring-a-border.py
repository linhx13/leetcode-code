from typing import List


class Solution:
    def colorBorder(self, grid: List[List[int]], r0: int, c0: int,
                    color: int) -> List[List[int]]:
        res = []
        for row in grid:
            res.append([x for x in row])
        self.dfs(grid, r0, c0, color, grid[r0][c0], res)
        return res

    def dfs(self, grid, r, c, color, ori_color, res):
        if r < 0 or r >= len(grid) or c < 0 or c >= len(grid[0]):
            return
        if grid[r][c] != ori_color:
            return
        dirs = [[0, -1], [0, 1], [-1, 0], [1, 0]]
        for d in dirs:
            rr, cc = r + d[0], c + d[1]
            if rr < 0 or rr >= len(grid) or cc < 0 or cc >= len(grid[0]):
                res[r][c] = color
            elif grid[rr][cc] != ori_color and grid[rr][cc] != -ori_color:
                res[r][c] = color
            grid[r][c] = -grid[r][c]
            self.dfs(grid, rr, cc, color, ori_color, res)
            grid[r][c] = -grid[r][c]
