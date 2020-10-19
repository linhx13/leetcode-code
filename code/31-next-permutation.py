from typing import List


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i = len(nums) - 1
        while i >= 0:
            if i + 1 < len(nums) and nums[i] < nums[i + 1]:
                break
            i -= 1
        if i < 0:
            nums.sort()
            return
        min_idx, j = -1, len(nums) - 1
        while j > i:
            if nums[j] > nums[i]:
                if min_idx == -1 or nums[j] < nums[min_idx]:
                    min_idx = j
            j -= 1
        print(i, min_idx)
        nums[i], nums[min_idx] = nums[min_idx], nums[i]
        nums[i + 1:] = sorted(nums[i + 1:])


if __name__ == "__main__":
    nums = [2, 3, 1]
    sol = Solution()
    sol.nextPermutation(nums)
    print(nums)
