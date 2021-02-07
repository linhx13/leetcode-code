from typing import List
import collections


class Solution:
    def countGoodRectangles(self, rectangles: List[List[int]]) -> int:
        lens = list(map(min, rectangles))
        max_len = max(lens)
        return len([x for x in lens if x == max_len])
