class Solution:
    def soupServings(self, N: int) -> float:
        if N >= 5000:
            return 1.0
        memo = {}
        return self.helper(N, N, memo)

    def helper(self, A, B, memo):
        if A <= 0 and B > 0:
            memo[(A, B)] = 1.0
        elif A <= 0 and B <= 0:
            memo[(A, B)] = 0.5
        elif A > 0 and B <= 0:
            memo[(A, B)] = 0.0
        else:
            if (A, B) in memo:
                return memo[(A, B)]
            p1 = self.helper(A - 100, B, memo)
            p2 = self.helper(A - 75, B - 25, memo)
            p3 = self.helper(A - 50, B - 50, memo)
            p4 = self.helper(A - 25, B - 75, memo)
            memo[(A, B)] = 0.25 * (p1 + p2 + p3 + p4)
        return memo[(A, B)]


if __name__ == "__main__":
    # N = 660295675
    N = 50
    print(Solution().soupServings(N))
