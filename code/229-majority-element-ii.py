from typing import List
from collections import Counter


class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        n = len(nums) // 3
        counter = Counter(nums)
        return [k for k, v in counter.items() if v > n]
