from typing import List


class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        i, j = 0, 0
        p = 1
        res = 0
        while j < len(nums):
            p *= nums[j]
            while p >= k and i <= j:
                p //= nums[i]
                i += 1
            if i <= j:
                res += j - i + 1
            j += 1
        return res


if __name__ == "__main__":
    nums = [20]
    k = 20
    print(Solution().numSubarrayProductLessThanK(nums, k))
