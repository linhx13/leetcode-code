from typing import List


class Solution:
    def numberOfArithmeticSlices(self, A: List[int]) -> int:
        dp = [0] * len(A)
        res = 0
        for i in range(2, len(dp)):
            if A[i] - A[i - 1] == A[i - 1] - A[i - 2]:
                dp[i] = 1 + dp[i - 1]
                res += dp[i]
        return res
