from typing import List


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [float("inf") for _ in range(amount + 1)]
        dp[0] = 0

        for i in range(1, len(coins) + 1):
            for j in range(coins[i - 1], amount + 1):
                if dp[j] > 1 + dp[j - coins[i - 1]]:
                    dp[j] = 1 + dp[j - coins[i - 1]]
        return dp[amount] if dp[amount] < float("inf") else -1
