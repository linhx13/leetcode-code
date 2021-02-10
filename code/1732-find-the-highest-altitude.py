from typing import List


class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        a = 0
        res = 0
        for x in gain:
            a += x
            res = max(res, a)
        return res
