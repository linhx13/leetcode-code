from typing import List


class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        arr = [p[0] for p in points]
        arr.sort()
        ans = 0
        for i in range(0, len(arr) - 1):
            ans = max(ans, arr[i + 1] - arr[i])
        return ans
