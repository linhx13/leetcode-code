from typing import List


class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        dd = {0: -1}
        s = 0
        res = 0
        for i, x in enumerate(nums):
            if x == 0:
                x = -1
            s += x
            if s in dd:
                res = max(res, i - dd[s])
            else:
                dd[s] = i
        return res
