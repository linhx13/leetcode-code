class Solution:
    def integerReplacement(self, n: int) -> int:
        memo = {}
        return self.helper(n, memo)

    def helper(self, n, memo):
        if n in memo:
            return memo[n]
        if n == 1:
            return 0
        elif n % 2 == 0:
            memo[n] = 1 + self.helper(n >> 1, memo)
        else:
            memo[n] = min(self.helper(n + 1, memo), self.helper(n - 1,
                                                                memo)) + 1
        return memo[n]
