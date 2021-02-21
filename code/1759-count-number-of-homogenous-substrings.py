import itertools


class Solution:
    def countHomogenous(self, s: str) -> int:
        mod = 10 ** 9 + 7
        res = 0
        for _, g in itertools.groupby(s):
            k = sum(1 for _ in g)
            res = (res + k * (k + 1) // 2) % mod
        return res
