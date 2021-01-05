from typing import List


"""
dp[i][j][k] 使用前i个str，得到少于m个0和n个1的最大子集数量

dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeros[i]][k-ones[j]] + 1)
"""


class Solution:
    def findMaxForm(self, strs: List[str], m: int, n: int) -> int:
        dp = [[0 for _ in range(n + 1)] for _ in range(m + 1)]

        for i in range(1, len(strs) + 1):
            w0, w1 = 0, 0
            for c in strs[i - 1]:
                if c == "0":
                    w0 += 1
                else:
                    w1 += 1

            for j in range(m, w0 - 1, -1):
                for k in range(n, w1 - 1, -1):
                    dp[j][k] = max(dp[j][k], dp[j - w0][k - w1] + 1)
        return dp[m][n]
