from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if not nums:
            return 0
        i, j = 0, 0
        cnt = 0
        last = nums[j]
        while i < len(nums) and j < len(nums):
            if last == nums[j]:
                cnt += 1
                if cnt <= 2:
                    nums[i] = last
                    i += 1
            else:
                last = nums[j]
                cnt = 1
                nums[i] = nums[j]
                i += 1
            j += 1
        return i


if __name__ == '__main__':
    # nums = [1, 1, 1, 2, 2, 3]
    # nums = [0, 0, 1, 1, 1, 1, 2, 3, 3]
    nums = []
    print(Solution().removeDuplicates(nums))
    print(nums)
