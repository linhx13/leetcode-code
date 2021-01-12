from typing import List


class Solution:
    def flipgame(self, fronts: List[int], backs: List[int]) -> int:
        same = {x for i, x in enumerate(fronts) if x == backs[i]}
        res = float("inf")
        for x in fronts:
            if x not in same:
                res = min(res, x)
        for x in backs:
            if x not in same:
                res = min(res, x)
        return res if res < float("inf") else 0
