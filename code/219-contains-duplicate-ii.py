from typing import List


class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        d = {}
        for i, x in enumerate(nums):
            j = d.get(x, -1)
            if j >= 0 and abs(i - j) <= k:
                return True
            d[x] = i
        return False
