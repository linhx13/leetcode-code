from typing import List


class Solution:
    def cherryPickup(self, grid: List[List[int]]) -> int:
        n = len(grid)
        dp = [[float("-inf")] * n for _ in range(n)]
        dp[0][0] = grid[0][0]
        for t in range(1, 2 * n - 1):
            dp2 = [[float("-inf")] * n for _ in range(n)]
            for i in range(max(0, t - (n - 1)), min(n - 1, t) + 1):
                for j in range(max(0, t - (n - 1)), min(n - 1, t) + 1):
                    if grid[i][t - i] == -1 or grid[j][t - j] == -1:
                        continue
                    val = grid[i][t - i]
                    if i != j:
                        val += grid[j][t - j]
                    dp2[i][j] = max(
                        dp[pi][pj] + val
                        for pi in (i - 1, i)
                        for pj in (j - 1, j)
                        if pi >= 0 and pj >= 0
                    )
            dp = dp2
        return max(0, dp[n - 1][n - 1])
