from typing import List


class Solution:
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        d = set()
        i, j = 0, 0
        s = 0
        res = 0
        n = len(nums)
        while i < n and j < n:
            if nums[j] not in d:
                d.add(nums[j])
                s += nums[j]
                res = max(res, s)
                j += 1
            else:
                s -= nums[i]
                d.remove(nums[i])
                i += 1
        return res
