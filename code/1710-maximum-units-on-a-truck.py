from typing import List


class Solution:
    def maximumUnits(self, boxTypes: List[List[int]], truckSize: int) -> int:
        boxTypes.sort(key=lambda b: -b[1])
        res = 0
        for x, y in boxTypes:
            res += min(x, truckSize) * y
            truckSize -= min(x, truckSize)
            if truckSize <= 0:
                break
        return res
