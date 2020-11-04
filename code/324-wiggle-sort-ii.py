from typing import List


class Solution:
    def wiggleSort(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        nums.sort()
        n = len(nums)
        if n % 2 == 0:
            a = nums[0:n // 2]
            a.reverse()
            nums[0:n // 2] = a
        else:
            a = nums[0:n // 2 + 1]
            a.reverse()
            nums[0:n // 2 + 1] = a
        left, right = 0, n - 1
        ans = []
        while left <= right:
            ans.append(nums[left])
            if left != right:
                ans.append(nums[right])
            left += 1
            right -= 1
        for i in range(n):
            nums[i] = ans[i]
