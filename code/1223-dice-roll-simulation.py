from typing import List


class Solution:
    def dieSimulator(self, n: int, rollMax: List[int]) -> int:
        kmod = 10**9 + 7
        dp = [[[0 for _ in range(15 + 1)] for _ in range(6)]
              for _ in range(n + 1)]
        for j in range(6):
            dp[1][j][1] = 1
        for i in range(2, n + 1):
            for j in range(6):
                for p in range(6):
                    for k in range(1, rollMax[p] + 1):
                        if p != j:
                            dp[i][j][1] = (dp[i][j][1] +
                                           dp[i - 1][p][k]) % kmod
                        elif k < rollMax[p]:
                            dp[i][j][k + 1] = dp[i - 1][j][k]
        res = 0
        for j in range(6):
            for k in range(1, rollMax[j] + 1):
                res = (res + dp[n][j][k]) % kmod
        return res


if __name__ == "__main__":
    n = 2
    rollMax = [1, 1, 2, 2, 2, 3]
    print(Solution().dieSimulator(n, rollMax))
