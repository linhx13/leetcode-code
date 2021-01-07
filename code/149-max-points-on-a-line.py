from typing import List
import sys
import math


class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        def calc_slope(x1, y1, x2, y2):
            dx, dy = x1 - x2, y1 - y2
            if dx == 0:
                return (0, 0)
            elif dy == 0:
                return (sys.maxsize, sys.maxsize)
            elif dx < 0:
                dx, dy = -dx, -dy
            gcd = math.gcd(dx, dy)
            return (dx // gcd, dy // gcd)

        def max_point_i(i):
            def add_line(i, j, count, dups):
                x1, y1 = points[i]
                x2, y2 = points[j]
                if x1 == x2 and y1 == y2:
                    dups += 1
                elif y1 == y2:
                    nonlocal horizon_lines
                    horizon_lines += 1
                    count = max(count, horizon_lines)
                else:
                    slop = calc_slope(x1, y1, x2, y2)
                    lines[slop] = lines.get(slop, 1) + 1
                    count = max(lines[slop], count)
                return count, dups

            lines, horizon_lines = {}, 1
            count = 1
            dups = 0
            nonlocal n
            for j in range(i + 1, n):
                count, dups = add_line(i, j, count, dups)
            return count + dups

        n = len(points)
        if n < 3:
            return n
        res = 1
        for i in range(n - 1):
            res = max(res, max_point_i(i))
        return res
