from typing import List


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        sums = []
        for i in range(len(nums)):
            s = nums[i]
            sums.append(s)
            for j in range(i + 1, len(nums)):
                s += nums[j]
                sums.append(s)
        sums.sort()
        res = 0
        for i in range(left - 1, right):
            res = (res + sums[i]) % 1000000007
        return res
