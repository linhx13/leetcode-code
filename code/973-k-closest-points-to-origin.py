from typing import List


class Solution:
    def kClosest(self, points: List[List[int]], K: int) -> List[List[int]]:
        def key_func(p):
            return p[0] * p[0] + p[1] * p[1]

        points.sort(key=key_func)
        return points[:K]
