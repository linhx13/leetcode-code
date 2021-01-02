from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        def helper(arr):
            if len(arr) == 0:
                return 0
            if len(arr) == 1:
                return arr[0]
            dp = [0] * len(arr)
            dp[0] = arr[0]
            dp[1] = max(arr[0], arr[1])
            for i in range(2, len(arr)):
                dp[i] = max(arr[i] + dp[i - 2], dp[i - 1])
            return dp[-1]

        if len(nums) == 0:
            return 0
        if len(nums) == 1:
            return nums[0]
        return max(helper(nums[: len(nums) - 1]), helper(nums[1:]))
