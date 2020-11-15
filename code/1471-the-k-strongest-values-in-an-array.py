from typing import List


class Solution:
    def getStrongest(self, arr: List[int], k: int) -> List[int]:
        arr.sort()
        idx = (len(arr) - 1) // 2
        m = arr[idx]
        arr.sort(key=lambda x: (-abs(x - m), -x))
        return arr[:k]
