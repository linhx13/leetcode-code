class Solution:
    def countVowelStrings(self, n: int) -> int:
        chars = "aeiou"
        memo = {}
        return self.helper(n, 0, chars, memo)

    def helper(self, n, last_idx, chars, memo):
        if n == 0:
            return 0
        if n == 1:
            return len(chars) - last_idx
        if (n, last_idx) in memo:
            return memo[(n, last_idx)]
        ans = 0
        for i in range(last_idx, len(chars)):
            ans += self.helper(n - 1, i, chars, memo)
        memo[(n, last_idx)] = ans
        return ans


if __name__ == "__main__":
    n = 50
    print(Solution().countVowelStrings(n))
