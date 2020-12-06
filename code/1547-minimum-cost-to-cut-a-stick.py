from typing import List


class Solution:
    def minCost(self, n: int, cuts: List[int]) -> int:
        m = len(cuts)
        cuts = [0] + sorted(cuts) + [n]
        dp = [[0] * (m + 1) for _ in range(m + 1)]

        for l in range(2, m + 2):
            for i in range(0, m + 2 - l):
                j = i + l - 1
                dp[i][j] = 999999999
                for k in range(i, j):
                    dp[i][j] = min(
                        dp[i][j],
                        dp[i][k] + dp[k + 1][j] + cuts[j + 1] - cuts[i])
        return dp[0][m]
