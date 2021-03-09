import functools


class Solution:
    def getKth(self, lo: int, hi: int, k: int) -> int:
        @functools.lru_cache(None)
        def power(x):
            if x == 1:
                return 0
            if x & 1 == 0:
                return power(x >> 1) + 1
            else:
                return power(3 * x + 1) + 1

        ps = [(power(x), x) for x in range(lo, hi + 1)]
        ps.sort()
        return ps[k - 1][1]
