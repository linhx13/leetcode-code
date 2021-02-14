from typing import List
import collections


class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        counter = collections.Counter(nums)
        return sum(k for k, v in counter.items() if v == 1)
