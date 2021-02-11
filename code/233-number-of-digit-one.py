class Solution:
    def countDigitOne(self, n: int) -> int:
        res = 0
        i = 1
        while i <= n:
            d = i * 10
            res += (n // d) * i + min(max(n % d - i + 1, 0), i)
            i = d
        return res
