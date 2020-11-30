from typing import List
from bisect import bisect_left


class Solution:
    def findRightInterval(self, intervals: List[List[int]]) -> List[int]:
        s = []
        sd = {}
        e = []
        cnt = 0
        for i in intervals:
            s.append(i[0])
            sd[i[0]] = cnt
            cnt += 1
            e.append(i[1])
        s.sort()
        n = len(s)
        res = [-1] * n
        for i in range(len(e)):
            j = bisect_left(s, e[i])
            if j < n:
                res[i] = sd[s[j]]
        return res
