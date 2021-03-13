from typing import List
import heapq


class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        n = len(matrix)
        h = []
        for j in range(n):
            heapq.heappush(h, (matrix[0][j], 0, j))
        for i in range(k - 1):
            v, r, c = heapq.heappop(h)
            if r == n - 1:
                continue
            heapq.heappush(h, (matrix[r + 1][c], r + 1, c))
        return h[0][0]
