from typing import List
from collections import Counter


class Solution:
    def frequencySort(self, nums: List[int]) -> List[int]:
        freqs = Counter(nums)
        return sorted(nums, key=lambda x: (freqs[x], -x))
