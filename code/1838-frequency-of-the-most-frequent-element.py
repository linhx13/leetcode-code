from typing import List


class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        nums.sort()
        res = 0
        i, j = 0, 0
        cur = 0
        while j < len(nums):
            cur += nums[j]
            while cur + k < nums[j] * (j - i + 1):
                cur -= nums[i]
                i += 1
            res = max(res, j - i + 1)
            j += 1
        return res
