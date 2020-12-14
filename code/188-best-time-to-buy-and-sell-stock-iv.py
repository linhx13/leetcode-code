from typing import List


class Solution:
    def maxProfit(self, k: int, prices: List[int]) -> int:
        K = k
        if len(prices) == 0:
            return 0
        dp = [[0 for _ in range(len(prices))] for _ in range(K + 1)]
        for k in range(1, K + 1):
            m = prices[0]
            for i in range(1, len(prices)):
                m = min(m, prices[i] - dp[k - 1][i - 1])
                dp[k][i] = max(dp[k][i - 1], prices[i] - m)
        return dp[K][len(prices) - 1]
