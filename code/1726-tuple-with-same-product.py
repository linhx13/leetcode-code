from typing import List
import collections


class Solution:
    def tupleSameProduct(self, nums: List[int]) -> int:
        counter = collections.Counter(
            nums[i] * nums[j]
            for i in range(len(nums))
            for j in range(i + 1, len(nums))
        )

        res = sum((v * (v - 1) << 2) for v in counter.values())
        return res
