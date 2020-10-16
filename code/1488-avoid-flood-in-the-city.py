from typing import List
from bisect import bisect_right


class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        full = {}
        i = 0
        dry = []
        res = [-1 for _ in range(len(rains))]
        for i in range(len(rains)):
            if rains[i] == 0:
                dry.append(i)
                continue
            if not rains[i] in full:
                full[rains[i]] = i
                continue
            idx = bisect_right(dry, full[rains[i]])
            if idx == len(dry):
                return []
            res[dry[idx]] = rains[i]
            full[rains[i]] = i
            dry.pop(idx)
        for x in dry:
            res[x] = 1
        return res
