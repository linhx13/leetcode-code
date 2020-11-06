from typing import List


class Solution:
    def maxAbsValExpr(self, arr1: List[int], arr2: List[int]) -> int:
        return max(
            self.helper(sign1, sign2, arr1, arr2)
            for sign1, sign2 in [(1, 1), (-1, 1), (1, -1), (-1, -1)])

    def helper(self, sign1, sign2, arr1, arr2):
        minn = float('inf')
        maxx = float('-inf')
        for i in range(len(arr1)):
            x = sign1 * arr1[i] + sign2 * arr2[i] + i
            minn = min(minn, x)
            maxx = max(maxx, x)
        return maxx - minn
