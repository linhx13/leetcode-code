class Solution:
    def knightDialer(self, n: int) -> int:
        moves = {
            1: {6, 8},
            2: {7, 9},
            3: {4, 8},
            4: {0, 3, 9},
            5: {},
            6: {0, 1, 7},
            7: {2, 6},
            8: {1, 3},
            9: {2, 4},
            0: {4, 6},
        }
        MOD = 10 ** 9 + 7

        dp = [1] * 10

        for k in range(2, n + 1):
            dp2 = [0] * 10
            for cur in range(0, 10):
                for x in moves[cur]:
                    dp2[cur] += dp[x]
                    dp2[cur] %= MOD
            dp = dp2
        return sum(dp) % MOD


if __name__ == "__main__":
    n = 4
    print(Solution().knightDialer(n))
