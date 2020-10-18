from math import sqrt


class Solution:
    def numSquares(self, n: int) -> int:
        dp = [0] * (n + 1)
        sqrs = []
        x = 1
        i = 1
        while x <= n:
            sqrs.append(x)
            dp[x] = 1
            i += 1
            x = i * i
        return self.helper(n, dp, sqrs)

    def helper(self, r, dp, sqrs):
        if r == 0:
            return 0
        if dp[r] != 0:
            return dp[r]
        dp[r] = r
        for i in sqrs:
            if r >= i:
                cur = self.helper(r % i, dp, sqrs) + r // i
                dp[r] = min(dp[r], cur)
            else:
                break
        return dp[r]


if __name__ == "__main__":
    n = 7168
    sol = Solution()
    print(sol.numSquares(n))
