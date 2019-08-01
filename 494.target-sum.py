#
# @lc app=leetcode id=494 lang=python
#
# [494] Target Sum
#
# https://leetcode.com/problems/target-sum/description/
#
# algorithms
# Medium (44.79%)
# Total Accepted:    85.3K
# Total Submissions: 190.3K
# Testcase Example:  '[1,1,1,1,1]\n3'
#
#
# You are given a list of non-negative integers, a1, a2, ..., an, and a target,
# S. Now you have 2 symbols + and -. For each integer, you should choose one
# from + and - as its new symbol.
# ⁠
#
# Find out how many ways to assign symbols to make sum of integers equal to
# target S.
#
#
# Example 1:
#
# Input: nums is [1, 1, 1, 1, 1], S is 3.
# Output: 5
# Explanation:
#
# -1+1+1+1+1 = 3
# +1-1+1+1+1 = 3
# +1+1-1+1+1 = 3
# +1+1+1-1+1 = 3
# +1+1+1+1-1 = 3
#
# There are 5 ways to assign symbols to make the sum of nums be target 3.
#
#
#
# Note:
#
# The length of the given array is positive and will not exceed 20.
# The sum of elements in the given array will not exceed 1000.
# Your output answer is guaranteed to be fitted in a 32-bit integer.
#
#
#

import collections

class Solution(object):
    def __init__(self):
        self.result = 0

    def findTargetSumWays(self, nums, S):
        """
        :type nums: List[int]
        :type S: int
        :rtype: int
        """
        _len = len(nums)
        dp = [collections.defaultdict(int) for _ in range(_len + 1)]
        dp[0][0] = 1
        for i, num in enumerate(nums):
            for sum, cnt in dp[i].items():
                dp[i + 1][sum + num] += cnt
                dp[i + 1][sum - num] += cnt
        return dp[_len][S]

    # def findTargetSumWays(self, nums, S):
    #     """
    #     :type nums: List[int]
    #     :type S: int
    #     :rtype: int
    #     """
    #     self.nums = nums
    #     self.s = S

    #     return self.dfs(0,0)

    # def dfs(self, index, s):
    #     if index == len(self.nums):
    #         if s == self.s:
    #             return 1
    #         else:
    #             return 0

    #     return self.dfs(index + 1, s + self.nums[index]) + self.dfs(
    #         index + 1, s - self.nums[index])
