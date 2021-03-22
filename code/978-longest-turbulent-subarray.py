from typing import List


class Solution:
    def maxTurbulenceSize(self, arr: List[int]) -> int:
        N = len(arr)
        res = 1
        k = 0

        def cmp(x, y):
            if x < y:
                return -1
            elif x > y:
                return 1
            else:
                return 0

        for i in range(1, N):
            c = cmp(arr[i - 1], arr[i])
            if c == 0:
                k = i
            elif i == N - 1 or c * cmp(arr[i], arr[i + 1]) != -1:
                res = max(res, i - k + 1)
                k = i
        return res
