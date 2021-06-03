from typing import List


class Solution:
    def maxArea(
        self,
        h: int,
        w: int,
        horizontalCuts: List[int],
        verticalCuts: List[int],
    ) -> int:
        horizontalCuts.sort()
        verticalCuts.sort()
        max_h = max(horizontalCuts[0], h - horizontalCuts[-1])
        max_v = max(verticalCuts[0], w - verticalCuts[-1])
        for i in range(0, len(horizontalCuts) - 1):
            max_h = max(max_h, horizontalCuts[i + 1] - horizontalCuts[i])
        for i in range(0, len(verticalCuts) - 1):
            max_v = max(max_v, verticalCuts[i + 1] - verticalCuts[i])
        return max_h * max_v % 1000000007
