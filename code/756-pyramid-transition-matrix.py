from typing import List
from collections import defaultdict
from itertools import product


class Solution:
    def pyramidTransition(self, bottom: str, allowed: List[str]) -> bool:
        dd = defaultdict(set)
        for x, y, z in allowed:
            dd[(x, y)].add(z)
        return self.dfs(bottom, dd)

    def dfs(self, cur, dd):
        if len(cur) == 1:
            return True
        for xx in product(*(dd[(x, y)] for x, y in zip(cur, cur[1:]))):
            if self.dfs(xx, dd):
                return True
        return False
