from typing import List


class Solution:
    def nearestValidPoint(
        self, x: int, y: int, points: List[List[int]]
    ) -> int:
        res, mind = -1, float("inf")
        for i, p in enumerate(points):
            if not (x == p[0] or y == p[1]):
                continue
            d = abs(p[0] - x) + abs(p[1] - y)
            if d < mind:
                res, mind = i, d
        return res
