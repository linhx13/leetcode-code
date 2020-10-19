from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        if len(nums) == 0:
            return []
        res: List[int] = [1 for _ in range(len(nums))]
        for i in range(1, len(nums)):
            res[i] = nums[i - 1] * res[i - 1]
        p = 1
        for i in range(len(nums) - 2, -1, -1):
            p *= nums[i + 1]
            res[i] *= p
        return res


if __name__ == "__main__":
    nums = [1, 2, 3, 4]
    sol = Solution()
    print(sol.productExceptSelf(nums))
