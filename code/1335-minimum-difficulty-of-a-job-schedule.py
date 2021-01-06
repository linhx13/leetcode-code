from typing import List


class Solution:
    def minDifficulty(self, jobDifficulty: List[int], d: int) -> int:
        n = len(jobDifficulty)
        dp = [[float("inf")] * n + [0] for i in range(d + 1)]
        for k in range(1, d + 1):
            for i in range(n - k + 1):
                maxd = 0
                for j in range(i, n - k + 1):
                    maxd = max(maxd, jobDifficulty[j])
                    dp[k][i] = min(dp[k][i], maxd + dp[k - 1][j + 1])
        return dp[d][0] if dp[d][0] < float("inf") else -1
