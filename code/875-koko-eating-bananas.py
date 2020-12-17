from typing import List


class Solution:
    def minEatingSpeed(self, piles: List[int], H: int) -> int:
        def possible(k):
            return sum((p - 1) // k + 1 for p in piles) <= H

        low, high = 1, max(piles)
        while low < high:
            mid = low + (high - low) // 2
            if not possible(mid):
                low = mid + 1
            else:
                high = mid
        return low
