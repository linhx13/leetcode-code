from typing import List


class Solution:
    def countServers(self, grid: List[List[int]]) -> int:
        rows = [0] * len(grid)
        cols = [0] * len(grid[0])
        for x, row in enumerate(grid):
            for y, v in enumerate(row):
                if v == 1:
                    rows[x] += 1
                    cols[y] += 1
        res = 0
        for x, row in enumerate(grid):
            for y, v in enumerate(row):
                if v == 1 and (rows[x] > 1 or cols[y] > 1):
                    res += 1
        return res


if __name__ == "__main__":
    # grid = [[1, 0], [0, 1]]
    # grid = [[1, 0], [1, 1]]
    grid = [[1, 1, 0, 0], [0, 0, 1, 0], [0, 0, 1, 0], [0, 0, 0, 1]]
    sol = Solution()
    print(sol.countServers(grid))
