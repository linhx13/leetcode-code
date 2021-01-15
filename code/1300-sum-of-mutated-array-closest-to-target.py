from typing import List


class Solution:
    def findBestValue(self, arr: List[int], target: int) -> int:
        s = sum(arr)
        arr.sort()
        idx = len(arr) - 1
        removed = 0
        while idx >= 0 and target < s + removed * arr[idx]:
            s -= arr[idx]
            removed += 1
            idx -= 1
        if removed == 0:
            return arr[-1]
        v = (target - s) // removed
        if abs(target - s - removed * v) <= abs(
            target - s - removed * (v + 1)
        ):
            return v
        else:
            return v + 1
