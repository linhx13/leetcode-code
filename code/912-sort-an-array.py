from typing import List


class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        if not nums:
            return nums
        self.quick_sort(nums, 0, len(nums) - 1)
        return nums

    def quick_sort(self, nums, start, end):
        if start >= end:
            return
        left, right = start, end
        mid = left + (right - left) // 2
        pivot = nums[mid]

        while left <= right:
            while left <= right and nums[left] < pivot:
                left += 1
            while left <= right and nums[right] > pivot:
                right -= 1
            if left <= right:
                nums[left], nums[right] = nums[right], nums[left]
                left += 1
                right -= 1
        self.quick_sort(nums, start, right)
        self.quick_sort(nums, left, end)


if __name__ == '__main__':
    nums = [5, 2, 3, 1]
    sol = Solution()
    print(sol.sortArray(nums))
