from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) == 0:
            return 0
        dp = [[0 for _ in range(len(prices))] for _ in range(3)]
        for k in range(1, 3):
            m = prices[0]
            for i in range(1, len(prices)):
                m = min(m, prices[i] - dp[k - 1][i - 1])
                dp[k][i] = max(dp[k][i - 1], prices[i] - m)
        return dp[2][len(prices) - 1]


if __name__ == "__main__":
    # prices = [3, 3, 5, 0, 0, 3, 1, 4]
    # prices = [1, 2, 3, 4, 5]
    # prices = [7, 6, 4, 3, 1]
    # prices = [1]
    # prices = [1, 2, 4, 2, 5, 7, 2, 4, 9, 0]
    prices = [14, 9, 10, 12, 4, 8, 1, 16]
    # prices = list(range(10000, 0, -1))
    print(Solution().maxProfit(prices))
