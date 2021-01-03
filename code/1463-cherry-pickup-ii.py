from typing import List


class Solution:
    def cherryPickup(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        dp = {}
        if 0 == cols - 1:
            dp[(0, 0, cols - 1)] = grid[0][0]
        else:
            dp[(0, 0, cols - 1)] = grid[0][0] + grid[0][cols - 1]
        ds = [
            (-1, -1),
            (-1, 0),
            (-1, 1),
            (0, -1),
            (0, 0),
            (0, 1),
            (1, -1),
            (1, 0),
            (1, 1),
        ]
        for i in range(1, rows):
            for j in range(cols):
                for k in range(cols):
                    for d in ds:
                        jj, kk = j + d[0], k + d[1]
                        if not (0 <= jj < cols and 0 <= kk < cols):
                            continue
                        if (i - 1, jj, kk) not in dp:
                            continue
                        val = dp[(i - 1, jj, kk)] + grid[i][j] + grid[i][k]
                        if j == k:
                            val -= grid[i][j]
                        key = (i, j, k)
                        if key not in dp:
                            dp[key] = val
                        else:
                            dp[key] = max(val, dp[key])
        res = 0
        for j in range(cols):
            for k in range(cols):
                key = (rows - 1, j, k)
                if key in dp:
                    res = max(res, dp[key])
        return res


if __name__ == "__main__":
    # grid = [[3, 1, 1], [2, 5, 1], [1, 5, 5], [2, 1, 1]]
    # grid = [
    #     [1, 0, 0, 0, 0, 0, 1],
    #     [2, 0, 0, 0, 0, 3, 0],
    #     [2, 0, 9, 0, 0, 0, 0],
    #     [0, 3, 0, 5, 4, 0, 0],
    #     [1, 0, 2, 3, 0, 0, 6],
    # ]
    # grid = [[1, 0, 0, 3], [0, 0, 0, 3], [0, 0, 3, 3], [9, 0, 3, 3]]
    # grid = [[1, 1], [1, 1]]
    # grid = [[1]]
    grid = [
        [0, 8, 7, 10, 9, 10, 0, 9, 6],
        [8, 7, 10, 8, 7, 4, 9, 6, 10],
        [8, 1, 1, 5, 1, 5, 5, 1, 2],
        [9, 4, 10, 8, 8, 1, 9, 5, 0],
        [4, 3, 6, 10, 9, 2, 4, 8, 10],
        # [7, 3, 2, 8, 3, 3, 5, 9, 8],
        # [1, 2, 6, 5, 6, 2, 0, 10, 0],
    ]

    print(Solution().cherryPickup(grid))
