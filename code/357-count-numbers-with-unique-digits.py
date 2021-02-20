class Solution:
    def countNumbersWithUniqueDigits(self, n: int) -> int:
        if n == 0:
            return 1
        res = 10
        uniq = 9
        avail = 9
        while n > 1 and avail > 0:
            uniq = uniq * avail
            res += uniq
            avail -= 1
            n -= 1
        return res
