from typing import List


class Solution:
    def maxAscendingSum(self, nums: List[int]) -> int:
        last = 0
        s = 0
        res = 0
        for x in nums:
            if x > last:
                s += x
                last = x
            else:
                res = max(res, s)
                s = x
                last = x
        res = max(res, s)
        return res
