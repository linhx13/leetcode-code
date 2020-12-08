from typing import List


class Solution:
    def minDifference(self, nums: List[int]) -> int:
        if len(nums) <= 3:
            return 0
        nums.sort()
        res = 2 ** 31 - 1
        for i in range(0, 4):
            j = 3 - i
            diff = abs(nums[i] - nums[len(nums) - 1 - j])
            res = min(diff, res)
        return res


if __name__ == "__main__":
    # nums = [5, 3, 2, 4]
    # nums = [1, 5, 0, 10, 14]
    # nums = [6, 6, 0, 1, 1, 4, 6]
    nums = [1, 5, 6, 14, 15]
    print(Solution().minDifference(nums))
