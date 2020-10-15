from typing import List


class Solution:
    def findTargetSumWays(self, nums: List[int], S: int) -> int:
        memo = {}

        def helper(idx, s):
            if idx == len(nums) and s == S:
                return 1
            elif idx == len(nums):
                return 0
            key = (idx, s)
            if key not in memo:
                memo[key] = helper(idx + 1, s + nums[idx]) + helper(
                    idx + 1, s - nums[idx])
            return memo[key]

        return helper(0, 0)


if __name__ == '__main__':
    nums = [1, 1, 1, 1, 1]
    S = 3
    sol = Solution()
    print(sol.findTargetSumWays(nums, S))
