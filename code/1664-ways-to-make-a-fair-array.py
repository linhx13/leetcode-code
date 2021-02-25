from typing import List


class Solution:
    def waysToMakeFair(self, nums: List[int]) -> int:
        even_sum = sum(nums[::2])
        odd_sum = sum(nums[1::2])
        cur_odd, cur_even = 0, 0
        res = 0
        for i, x in enumerate(nums):
            if i & 1:
                cur_odd += x
                if (
                    cur_even + odd_sum - cur_odd
                    == cur_odd - x + even_sum - cur_even
                ):
                    res += 1
            else:
                cur_even += x
                if (
                    cur_even - x + odd_sum - cur_odd
                    == cur_odd + even_sum - cur_even
                ):
                    res += 1
        return res
