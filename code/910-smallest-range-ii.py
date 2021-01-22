from typing import List


class Solution:
    def smallestRangeII(self, A: List[int], K: int) -> int:
        A.sort()
        mi, ma = A[0], A[-1]
        res = ma - mi
        for i in range(len(A) - 1):
            a, b = A[i], A[i + 1]
            res = min(res, max(ma - K, a + K) - min(mi + K, b - K))
        return res
