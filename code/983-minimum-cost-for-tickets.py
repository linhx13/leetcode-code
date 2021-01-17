from typing import List
import functools


class Solution:
    def mincostTickets(self, days: List[int], costs: List[int]) -> int:
        days = set(days)
        durations = [1, 7, 30]

        @functools.lru_cache(None)
        def dp(i):
            if i > 365:
                return 0
            elif i in days:
                return min(dp(i + d) + c for d, c in zip(durations, costs))
            else:
                return dp(i + 1)

        return dp(1)
