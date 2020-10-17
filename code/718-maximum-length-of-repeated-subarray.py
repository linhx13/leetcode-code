from typing import List


class Solution:
    def findLength(self, A: List[int], B: List[int]) -> int:
        res = 0
        m, n = len(A), len(B)
        if m == 0 or n == 0:
            return 0
        dp = [[0 for _ in range(m + 1)] for _ in range(n + 1)]
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if A[i - 1] == B[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1] + 1
                else:
                    dp[i][j] = 0
                res = max(res, dp[i][j])
        return res
