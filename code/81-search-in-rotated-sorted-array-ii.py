from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        return self.helper(nums, target, 0, len(nums))

    def helper(self, nums, target, low, high):
        while low < high:
            mid = low + (high - low) // 2
            # print(low, mid, high, nums[mid])
            if nums[mid] == target:
                return True
            elif nums[mid] == nums[low] or nums[mid] == nums[high - 1]:
                return (self.helper(nums, target, low, mid)
                        or self.helper(nums, target, mid + 1, high))
            elif nums[mid] > target:
                if nums[low] <= target or nums[high - 1] > nums[mid]:
                    high = mid
                else:
                    low = mid + 1
            elif nums[mid] < target:
                if nums[high - 1] >= target or nums[low] < nums[mid]:
                    low = mid + 1
                else:
                    high = mid
        return False


if __name__ == "__main__":
    nums = [2, 5, 6, 0, 0, 1, 2]
    target = 1
    # nums = [1, 1, 1, 3, 1]
    # target = 3
    print(Solution().search(nums, target))
