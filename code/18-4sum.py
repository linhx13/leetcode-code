from typing import List


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        def k_sum(nums, start, target, k):
            if start == len(
                    nums) or nums[start] * k > target or target > nums[-1] * k:
                return []
            if k == 2:
                return two_sum(nums, start, target)
            res = []
            for i in range(start, len(nums)):
                if i == start or nums[i - 1] != nums[i]:
                    for _, x in enumerate(
                            k_sum(nums, i + 1, target - nums[i], k - 1)):
                        res.append([nums[i]] + x)
            return res

        def two_sum(nums, start, target):
            res = []
            s = set()
            for i in range(start, len(nums)):
                if len(res) == 0 or res[-1][1] != nums[i]:
                    if target - nums[i] in s:
                        res.append([target - nums[i], nums[i]])
                s.add(nums[i])
            return res

        nums.sort()
        return k_sum(nums, 0, target, 4)
