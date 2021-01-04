from typing import List


class Solution:
    def wiggleMaxLength(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return len(nums)
        down, up = 1, 1
        for i in range(1, len(nums)):
            if nums[i] > nums[i - 1]:
                up = down + 1
            elif nums[i] < nums[i - 1]:
                down = up + 1
        return max(down, up)


if __name__ == "__main__":
    # nums = [1, 7, 4, 9, 2, 5]
    # nums = [1, 17, 5, 10, 13, 15, 10, 5, 16, 8]
    nums = [1, 2, 3, 4, 5, 6, 7, 8, 9]
    # nums = [1]
    print(Solution().wiggleMaxLength(nums))
