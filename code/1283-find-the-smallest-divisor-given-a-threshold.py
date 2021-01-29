from typing import List
from math import ceil


class Solution:
    def smallestDivisor(self, nums: List[int], threshold: int) -> int:
        def calc_sum(x):
            return sum(ceil(n / x) for n in nums)

        left, right = 1, max(nums)
        while left <= right:
            mid = (left + right) >> 1
            val = calc_sum(mid)
            if val > threshold:
                left = mid + 1
            else:
                right = mid - 1
        return left


if __name__ == "__main__":
    nums = [44, 22, 33, 11, 1]
    threshold = 5
    print(Solution().smallestDivisor(nums, threshold))
