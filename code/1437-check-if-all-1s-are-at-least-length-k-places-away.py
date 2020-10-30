from typing import List


class Solution:
    def kLengthApart(self, nums: List[int], k: int) -> bool:
        last_idx = -1
        for i, n in enumerate(nums):
            if n == 1:
                if last_idx != -1 and i - last_idx - 1 < k:
                    return False
                last_idx = i
        return True
