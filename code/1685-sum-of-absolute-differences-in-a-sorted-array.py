from typing import List


class Solution:
    def getSumAbsoluteDifferences(self, nums: List[int]) -> List[int]:
        res = [0 for _ in range(len(nums))]
        for i in range(1, len(nums)):
            res[0] += nums[i] - nums[0]
        for i in range(1, len(nums)):
            res[i] = (
                res[i - 1]
                + (nums[i] - nums[i - 1]) * i
                - (nums[i] - nums[i - 1]) * (len(nums) - i)
            )
        return res
