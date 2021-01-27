from typing import List
import heapq


class Solution:
    def restoreMatrix(
        self, rowSum: List[int], colSum: List[int]
    ) -> List[List[int]]:
        m, n = len(rowSum), len(colSum)
        res = [[0] * n for _ in range(m)]
        row_heap = []
        for i, s in enumerate(rowSum):
            heapq.heappush(row_heap, (s, i))
        col_heap = []
        for i, s in enumerate(colSum):
            heapq.heappush(col_heap, (s, i))
        while row_heap and col_heap:
            rs, ri = heapq.heappop(row_heap)
            cs, ci = heapq.heappop(col_heap)
            v = min(rs, cs)
            res[ri][ci] = v
            if rs - v > 0:
                heapq.heappush(row_heap, (rs - v, ri))
            if cs - v > 0:
                heapq.heappush(col_heap, (cs - v, ci))
        return res


if __name__ == "__main__":
    # rowSum = [3, 8]
    # colSum = [4, 7]
    # rowSum = [5, 7, 10]
    # colSum = [8, 6, 8]
    # rowSum = [14, 9]
    # colSum = [6, 9, 8]
    rowSum = [0]
    colSum = [1]
    print(Solution().restoreMatrix(rowSum, colSum))
