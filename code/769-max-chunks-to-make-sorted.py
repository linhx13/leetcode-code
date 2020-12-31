from typing import List


class Solution:
    def maxChunksToSorted(self, arr: List[int]) -> int:
        res = 0
        max_num = 0
        for i, x in enumerate(arr):
            max_num = max(max_num, x)
            if max_num == i:
                res += 1
        return res
