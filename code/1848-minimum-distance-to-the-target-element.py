from typing import List


class Solution:
    def getMinDistance(self, nums: List[int], target: int, start: int) -> int:
        res = len(nums)
        for i in range(start, -1, -1):
            if nums[i] == target:
                res = min(start - i, res)
                break
        for i in range(start, len(nums)):
            if nums[i] == target:
                res = min(i - start, res)
                break
        return res
