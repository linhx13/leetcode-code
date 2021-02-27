from typing import List


class Solution:
    def maxAbsoluteSum(self, nums: List[int]) -> int:
        min_sum, max_sum = 0, 0
        min_val, max_val = 0, 0
        for x in nums:
            max_val = max(0, max_val + x)
            max_sum = max(max_sum, max_val)
            min_val = min(0, min_val + x)
            min_sum = min(min_sum, min_val)
        return max(abs(min_sum), abs(max_sum))
