from typing import List


class Solution:
    def intervalIntersection(
        self, A: List[List[int]], B: List[List[int]]
    ) -> List[List[int]]:
        i, j = 0, 0
        m, n = len(A), len(B)
        res = []
        while i < m and j < n:
            if A[i][0] <= B[j][1] <= A[i][1]:
                res.append((max(A[i][0], B[j][0]), B[j][1]))
                j += 1
            elif B[j][0] <= A[i][1] <= B[j][1]:
                res.append((max(A[i][0], B[j][0]), A[i][1]))
                i += 1
            elif A[i][1] < B[j][0]:
                i += 1
            else:
                j += 1
        return res
