from typing import List


class Solution:
    def profitableSchemes(
        self, G: int, P: int, group: List[int], profit: List[int]
    ) -> int:
        MOD = 10 ** 9 + 7
        dp = [[0 for _ in range(P + 1)] for _ in range(G + 1)]
        for i in range(G + 1):
            dp[i][0] = 1

        for i in range(0, len(group)):
            for j in range(G, group[i] - 1, -1):
                for p in range(0, P + 1):
                    idx = min(profit[i] + p, P)
                    dp[j][idx] = (dp[j][idx] + dp[j - group[i]][p]) % MOD
        return dp[G][P]
