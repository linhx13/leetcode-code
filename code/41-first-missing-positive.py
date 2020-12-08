from typing import List


class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        if not nums:
            return 1
        n = len(nums)
        for i in range(n):
            cur = nums[i]
            while cur - 1 >= 0 and cur - 1 < n and cur != nums[cur - 1]:
                next = nums[cur - 1]
                nums[cur - 1] = cur
                cur = next
        for i in range(n):
            if nums[i] != i + 1:
                return i + 1
        return n + 1


if __name__ == "__main__":
    nums = [3, 4, -1, 1]
    print(Solution().firstMissingPositive(nums))
