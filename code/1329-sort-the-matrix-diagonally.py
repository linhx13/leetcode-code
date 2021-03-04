from typing import List
from collections import defaultdict


class Solution:
    def diagonalSort(self, mat: List[List[int]]) -> List[List[int]]:
        dd = defaultdict(list)
        m, n = len(mat), len(mat[0])
        for i in range(m):
            for j in range(n):
                dd[i - j].append(mat[i][j])
        res = [[0] * n for _ in range(m)]
        for k, v in dd.items():
            i = k if k >= 0 else 0
            for x in sorted(v):
                res[i][i - k] = x
                i += 1
        return res
