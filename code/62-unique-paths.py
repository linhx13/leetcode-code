class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        dp = [[-1 for _ in range(n)] for _ in range(m)]
        dp[m - 1][n - 1] = 1

        def dfs(i, j):
            if not (0 <= i < m and 0 <= j < n):
                return 0
            if dp[i][j] != -1:
                return dp[i][j]
            dp[i][j] = dfs(i + 1, j) + dfs(i, j + 1)
            return dp[i][j]

        return dfs(0, 0)
