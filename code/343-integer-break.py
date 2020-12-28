class Solution:
    def integerBreak(self, n: int) -> int:
        if n == 2:
            return 1
        if n == 3:
            return 2
        num_3 = n // 3
        r = n % 3
        if r == 1:
            r = 4
            num_3 -= 1
        elif r == 0:
            r = 1
        return int(3 ** num_3) * r
