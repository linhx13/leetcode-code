class Solution:
    def bulbSwitch(self, n: int) -> int:
        res = 0
        i = 1
        while i * i <= n:
            res += 1
            i += 1
        return res
