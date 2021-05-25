from typing import List


class Solution:
    def maximumPopulation(self, logs: List[List[int]]) -> int:
        res, k = 0, 0
        years = set()
        for x in logs:
            years.add(x[0])
            years.add(x[1])
        for year in sorted(years):
            cnt = sum(1 for x in logs if x[0] <= year < x[1])
            if cnt > res:
                res = cnt
                k = year
        return k
