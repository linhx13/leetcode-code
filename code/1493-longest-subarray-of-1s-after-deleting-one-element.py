from typing import List


class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        cnts = {}
        c = 0
        for i in range(len(nums)):
            if nums[i] == 0:
                cnts[i] = c
                c = 0
            else:
                c += 1
        if len(cnts) == 0:
            return len(nums) - 1
        res = 0
        c = 0
        for i in range(len(nums) - 1, -1, -1):
            if nums[i] == 0:
                res = max(res, cnts[i] + c)
                c = 0
            else:
                c += 1
        return res


if __name__ == "__main__":
    # nums = [1, 1, 0, 1]
    # nums = [0, 1, 1, 1, 0, 1, 1, 0, 1]
    # nums = [1, 1, 1]
    # nums = [0]
    nums = [1, 1, 0, 0, 1, 1, 1, 0, 1]
    print(Solution().longestSubarray(nums))
