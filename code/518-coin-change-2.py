from typing import List


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        if amount == 0:
            return 1
        if not coins:
            return 0
        dp = [0] * (amount + 1)
        for c in coins:
            for x in range(c, amount + 1):
                if x - c == 0:
                    dp[x] += 1
                else:
                    dp[x] += dp[x - c]
        return dp[amount]
