class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == -(2 ** 31) and divisor == -1:
            return 2 ** 31 - 1
        sign = -1 if (int(dividend > 0) ^ int(divisor > 0)) else 1
        a, b = abs(dividend), abs(divisor)
        res = 0
        while a >= b:
            t = b
            m = 1
            while (t << 1) <= a:
                t <<= 1
                m <<= 1
            a -= t
            res += m
        return sign * res
