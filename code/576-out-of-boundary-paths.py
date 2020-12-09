class Solution:
    def findPaths(self, m: int, n: int, N: int, i: int, j: int) -> int:
        mod = 1000000007
        memo = {}

        def dfs(i, j, k):
            if i == m or j == n or i < 0 or j < 0:
                return 1
            if k == 0:
                return 0
            key = (i, j, k)
            if memo.get(key, -1) >= 0:
                return memo[key]
            memo[key] = (
                dfs(i - 1, j, k - 1)
                + dfs(i + 1, j, k - 1)
                + dfs(i, j - 1, k - 1)
                + dfs(i, j + 1, k - 1)
            ) % mod
            return memo[key]

        return dfs(i, j, N)
