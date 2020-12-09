from typing import List
import heapq


class Solution:
    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        if not heights:
            return 0

        r, c = len(heights), len(heights[0])
        h = [(0, 0, 0)]
        visited = set()
        dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        res = 0

        while h:
            d, x, y = heapq.heappop(h)
            res = max(res, d)
            if x == r - 1 and y == c - 1:
                return res
            visited.add((x, y))
            for dd in dirs:
                nx, ny = x + dd[0], y + dd[1]
                if 0 <= nx < r and 0 <= ny < c and (nx, ny) not in visited:
                    nd = abs(heights[nx][ny] - heights[x][y])
                    heapq.heappush(h, (nd, nx, ny))
        return -1


if __name__ == "__main__":
    heights = [[1, 2, 2], [3, 8, 2], [5, 3, 5]]
    print(Solution().minimumEffortPath(heights))
