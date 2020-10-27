from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        low, high = 0, len(nums) - 1
        target = len(nums) - k
        while low < high:
            idx = self.partition(nums, low, high)
            if idx == target:
                return nums[idx]
            elif idx > target:
                high = idx - 1
            else:
                low = idx + 1
        return nums[high]

    def partition(self, nums, low, high):
        mid = low + (high - low) // 2
        val = nums[mid]
        nums[mid], nums[high] = nums[high], nums[mid]
        idx = low
        for i in range(low, high):
            if nums[i] < val:
                nums[idx], nums[i] = nums[i], nums[idx]
                idx += 1
        nums[idx], nums[high] = nums[high], nums[idx]
        return idx
