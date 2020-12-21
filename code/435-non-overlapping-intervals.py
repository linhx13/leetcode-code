from typing import List


class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        if not intervals:
            return 0
        intervals.sort(key=lambda x: x[1])
        cur_end = intervals[0][1]
        res = 0
        for x in intervals[1:]:
            if x[0] < cur_end:
                res += 1
                cur_end = min(cur_end, x[1])
            else:
                cur_end = x[1]
        return res
