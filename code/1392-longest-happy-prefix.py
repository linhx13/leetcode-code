class Solution:
    def longestPrefix(self, s: str) -> str:
        n = len(s)
        dp = [0] * n
        k = -1
        for i in range(1, n):
            while k >= 0 and s[i] != s[k + 1]:
                k = dp[k] - 1
            if s[i] == s[k + 1]:
                k += 1
            dp[i] = k + 1
        return s[0:dp[n - 1]]
