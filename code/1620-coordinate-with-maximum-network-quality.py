from typing import List
import math


class Solution:
    def bestCoordinate(self, towers: List[List[int]],
                       radius: int) -> List[int]:
        def dist(x1, y1, x2, y2):
            return math.sqrt((x1 - x2)**2 + (y1 - y2)**2)

        def quality(x, y):
            q = 0
            for t in towers:
                d = dist(x, y, t[0], t[1])
                if d <= radius:
                    q += int(t[2] / (1 + d))
            return q

        ans = (0, 0)
        ans_q = 0
        seen = set()
        for t in towers:
            for dx in range(-radius, radius + 1):
                for dy in range(-radius, radius + 1):
                    x, y = t[0] + dx, t[1] + dy
                    if x < 0 or y < 0:
                        continue
                    if (x, y) in seen:
                        continue
                    seen.add((x, y))
                    q = quality(x, y)
                    if q > ans_q:
                        ans_q = q
                        ans = (x, y)
                    elif q == ans_q:
                        if (x, y) < ans:
                            ans = (x, y)

        return ans


if __name__ == "__main__":
    # towers = [[1, 2, 5], [2, 1, 7], [3, 1, 9]]
    # radius = 2
    towers = [[42, 0, 0]]
    radius = 7
    print(Solution().bestCoordinate(towers, radius))
